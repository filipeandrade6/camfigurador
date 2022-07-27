package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	dac "github.com/xinsnake/go-http-digest-auth-client"
)

type CameraInfo struct {
	Fabricante  string
	Modelo      string
	NumeroSerie string
	MAC         string
	VersaoFW    string
}

func IdentificadorDeModelo(ip, usuario, senha string) (CameraInfo, error) {
	var c CameraInfo

	t := dac.NewTransport(usuario, senha)
	t.HTTPClient = &http.Client{Timeout: 2 * time.Second}

	_, status, err := Requisitador(t, fmt.Sprintf("http://%s/cgi-bin/magicBox.cgi?action=getSystemInfo", ip))
	if err != nil {
		return c, err
	}

	switch status {
	case 401:

	case 200:
		c.Fabricante = "dahua"

		// Parse Modelo --------------------------------
		body, status, err := Requisitador(t, fmt.Sprintf("http://%s/cgi-bin/magicBox.cgi?action=getDeviceType", ip)) // type=IPC-HFW5242E-ZE-MF
		if err != nil || status != 200 {
			return c, fmt.Errorf("status: %d, error: %w", status, err)
		}

		s := strings.Split(body, "=")
		c.Modelo = s[1]

		// Parse Serial Number -------------------------
		body, status, err = Requisitador(t, fmt.Sprintf("https://%s/cgi-bin/magicBox.cgi?action=getSerialNo", ip)) // sn=6G06324PAG4549F
		if err != nil || status != 200 {
			return c, fmt.Errorf("status: %d, error: %w", status, err)
		}

		s = strings.Split(body, "=")
		c.NumeroSerie = s[1]

		// Parse MAC Number ----------------------------
		body, status, err = Requisitador(t, fmt.Sprintf("http://%s/cgi-bin/configManager.cgi?action=getConfig&name=Network", ip)) // ...table.Network.eth0.PhysicalAddress=bc:32:5f:22:bf:89...
		if err != nil || status != 200 {
			return c, fmt.Errorf("status: %d, error: %w", status, err)
		}

		// TODO mais preparacao (mais dados agrupados)

		s = strings.Split(body, "=")
		c.MAC = s[1]

		// Parse Firmware Number -----------------------
		body, status, err = Requisitador(t, fmt.Sprintf("http://%s/cgi-bin/magicBox.cgi?action=getSoftwareVersion", ip)) // version=2.212.0000.0.R,build:2013-11-14
		if err != nil || status != 200 {
			return c, fmt.Errorf("status: %d, error: %w", status, err)
		}

		s = strings.Split(body, "=")
		c.VersaoFW = s[1]

		return c, nil

	case 404:

	default:
	}

	return c, nil
}

func Requisitador(t dac.DigestTransport, url string) (string, int, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", 0, err
	}

	res, err := t.RoundTrip(req)
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return "", 0, err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", 0, err
	}

	return string(resBody), res.StatusCode, nil
}
