package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	dac "github.com/xinsnake/go-http-digest-auth-client"
)

type ConfigurationInfo struct {
	Manufacturer string
	Model        string
	serialNumber string
	MAC          string
	Software     string
}

const (
	camModel = iota
	serialNumber
	mac
	software
)

func (m *model) IdentificadorDeModelo() (ConfigurationInfo, error) {
	var c ConfigurationInfo

	urlDahua := fmt.Sprintf("http://%s/cgi-bin/magicBox.cgi?action=getSystemInfo", m.deviceIP)
	urlAxis := fmt.Sprintf("http://%s/axis-cgi/param.cgi?action=list&group=Brand.ProdShortName", m.deviceIP)

	_, statusCode, err := Requisitador(m.httpTransport, urlDahua)
	if err != nil {
		return c, err
	}

	switch statusCode {
	case 401:
		return c, errors.New("erro na autenticação - verificar usuário e senha")
	case 200:
		c.Manufacturer = "dahua"
	case 404:
		_, statusCode, err = Requisitador(m.httpTransport, urlAxis)
		if err != nil {
			return c, err
		}

		switch statusCode {
		case 401:
			return c, errors.New("erro na autenticação - verificar usuário e senha")
		case 200:
			c.Manufacturer = "axis"
		case 404:
			return c, errors.New("erro ao identificar o dispositivo")
		default:
			return c, fmt.Errorf("status: %d, error: status desconhecido", statusCode)
		}

	default:
		return c, fmt.Errorf("status: %d, error: status desconhecido", statusCode)
	}

	urls := getURLs(c.Manufacturer)

	for k, v := range urls {
		body, status, err := Requisitador(m.httpTransport, fmt.Sprintf("http://%s%s", m.deviceIP, v))
		if err != nil || status != 200 {
			return c, fmt.Errorf("status: %d, error: %w", status, err)
		}
		s := strings.Split(body, "=")
		result := strings.TrimSuffix(strings.ReplaceAll(strings.ToUpper(s[1]), ":", ""), "\n")

		switch k {
		case camModel:
			c.Model = result
		case serialNumber:
			c.serialNumber = result
		case mac:
			c.MAC = result
		case software:
			c.Software = result
		}
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

func getURLs(manufacturer string) map[int]string {
	switch manufacturer {
	case "axis":
		return map[int]string{
			camModel:     "/axis-cgi/param.cgi?action=list&group=Brand.ProdShortName",
			serialNumber: "/axis-cgi/param.cgi?action=list&group=Properties.System.SerialNumber",
			mac:          "/axis-cgi/param.cgi?action=list&group=Network.eth0.MACAddress",
			software:     "/axis-cgi/param.cgi?action=list&group=Properties.Firmware.Version",
		}

	case "dahua":
		return map[int]string{
			camModel:     "/cgi-bin/magicBox.cgi?action=getDeviceType",
			serialNumber: "/cgi-bin/magicBox.cgi?action=getSerialNo",
			mac:          "/cgi-bin/configManager.cgi?action=getConfig&name=Network.eth0.PhysicalAddress",
			software:     "/cgi-bin/magicBox.cgi?action=getSoftwareVersion",
		}

	}

	return nil
}
