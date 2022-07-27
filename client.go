package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
	case 200:
		c.Fabricante = "dahua"

		// Parse Modelo --------------------------------
		body, status, err := Requisitador(t, fmt.Sprintf("http://%s/cgi-bin/magicBox.cgi?action=getDeviceType", ip)) // type=IPC-HFW5242E-ZE-MF

		// Parse Serial Number -------------------------
		body, status, err = Requisitador(t, fmt.Sprintf("https://%s/cgi-bin/magicBox.cgi?action=getSerialNo", ip)) // sn=6G06324PAG4549F

		// Parse MAC Number ----------------------------
		body, status, err = Requisitador(t,

		// Parse Firmware Number -----------------------
		body, status, err = Requisitador(t, fmt.Sprintf("http://%s/cgi-bin/magicBox.cgi?action=getSoftwareVersion", ip)) // version=2.212.0000.0.R,build:2013-11-14


		return c, nil
	}
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
