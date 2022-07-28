package main

import (
	"errors"
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
		return c, errors.New("erro na autenticação - verificar usuário e senha")

	case 200:
		c.Fabricante = "dahua"

		// Parse Modelo --------------------------------
		body, status, err := Requisitador(t, fmt.Sprintf("http://%s/cgi-bin/magicBox.cgi?action=getDeviceType", ip))
		if err != nil || status != 200 {
			return c, fmt.Errorf("status: %d, error: %w", status, err)
		}

		s := strings.Split(body, "=")
		c.Modelo = strings.ToUpper(s[1])

		// Parse Serial Number -------------------------
		body, status, err = Requisitador(t, fmt.Sprintf("https://%s/cgi-bin/magicBox.cgi?action=getSerialNo", ip))
		if err != nil || status != 200 {
			return c, fmt.Errorf("status: %d, error: %w", status, err)
		}

		s = strings.Split(body, "=")
		c.NumeroSerie = strings.ToUpper(s[1])

		// Parse MAC Number ----------------------------
		body, status, err = Requisitador(t, fmt.Sprintf("http://%s/cgi-bin/configManager.cgi?action=getConfig&name=Network", ip))
		if err != nil || status != 200 {
			return c, fmt.Errorf("status: %d, error: %w", status, err)
		}

		s = strings.Split(body, "eth0.PhysicalAddress=")
		s = strings.Split(s[1], "=")
		c.MAC = strings.Trim(strings.ToUpper(s[0]), ":")

		// Parse Firmware Number -----------------------
		body, status, err = Requisitador(t, fmt.Sprintf("http://%s/cgi-bin/magicBox.cgi?action=getSoftwareVersion", ip))
		if err != nil || status != 200 {
			return c, fmt.Errorf("status: %d, error: %w", status, err)
		}

		s = strings.Split(body, "=")
		s = strings.Split(s[1], ",")
		c.VersaoFW = s[0]

		return c, nil

	case 404:
		// Parse Modelo --------------------------------
		body, status, err := Requisitador(t, fmt.Sprintf("http://%s/axis-cgi/param.cgi?action=list&group=Brand.ProdShortName", ip))
		if err != nil || status != 200 {
			return c, fmt.Errorf("status: %d, error: %w", status, err)
		}

		switch status {
		case 401:
			return c, errors.New("erro na autenticação - verificar usuário e senha")

		case 200:
			c.Fabricante = "axis"

			// Parse Modelo --------------------------------
			s := strings.Split(body, "=")
			c.Modelo = s[1]

			// Parse Serial Number -------------------------
			body, status, err = Requisitador(t, fmt.Sprintf("http://%s/axis-cgi/param.cgi?action=list&group=Properties.System.SerialNumber", ip))
			if err != nil || status != 200 {
				return c, fmt.Errorf("status: %d, error: %w", status, err)
			}
			s = strings.Split(body, "=")
			c.NumeroSerie = strings.ToUpper(s[1])

			// Parse MAC Number ----------------------------
			body, status, err = Requisitador(t, fmt.Sprintf("http://%s/axis-cgi/param.cgi?action=list&group=Network.eth0.MACAddress", ip))
			if err != nil || status != 200 {
				return c, fmt.Errorf("status: %d, error: %w", status, err)
			}
			s = strings.Split(body, "=")
			c.MAC = strings.Trim(strings.ToUpper(s[1]), ":")

			// Parse Firmware Number -----------------------
			body, status, err = Requisitador(t, fmt.Sprintf("http://%s/axis-cgi/param.cgi?action=list&group=Properties.Firmware.Version", ip))
			if err != nil || status != 200 {
				return c, fmt.Errorf("status: %d, error: %w", status, err)
			}
			s = strings.Split(body, "=")
			c.VersaoFW = strings.ToUpper(s[1])

		case 404:
			return c, errors.New("dispositivo desconhecido")
			// TODO testar novo VAPIX (tem que utilizar POST Method)

		default:
			return c, fmt.Errorf("status: %d, error: status desconhecido", status)

		}

	default:
		return c, fmt.Errorf("status: %d, error: status desconhecido", status)

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
