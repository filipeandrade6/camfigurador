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

	return nil // TODO verificar
}

func IdentificadorDeModelo(ip, usuario, senha string) (CameraInfo, error) {
	var c CameraInfo
	response := make(map[int]string)

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
		c.Manufacturer = "dahua"
		response = getURLs("dahua")

		// // Parse Modelo --------------------------------
		// body, status, err := Requisitador(t, fmt.Sprintf("http://%s/cgi-bin/magicBox.cgi?action=getDeviceType", ip))
		// if err != nil || status != 200 {
		// 	return c, fmt.Errorf("status: %d, error: %w", status, err)
		// }

		// s := strings.Split(body, "=")
		// c.Model = strings.ToUpper(s[1])
		// fmt.Println(c.Model)

		// // Parse Serial Number -------------------------
		// body, status, err = Requisitador(t, fmt.Sprintf("http://%s/cgi-bin/magicBox.cgi?action=getSerialNo", ip))
		// if err != nil || status != 200 {
		// 	return c, fmt.Errorf("status: %d, error: %w", status, err)
		// }

		// s = strings.Split(body, "=")
		// c.serialNumber = strings.ToUpper(s[1])

		// // Parse MAC Number ----------------------------
		// body, status, err = Requisitador(t, fmt.Sprintf("http://%s/cgi-bin/configManager.cgi?action=getConfig&name=Network.eth0.PhysicalAddress", ip))
		// if err != nil || status != 200 {
		// 	return c, fmt.Errorf("status: %d, error: %w", status, err)
		// }

		// s = strings.Split(body, "=")
		// c.MAC = strings.ReplaceAll(strings.ToUpper(s[1]), ":", "")

		// // Parse Firmware Number -----------------------
		// body, status, err = Requisitador(t, fmt.Sprintf("http://%s/cgi-bin/magicBox.cgi?action=getSoftwareVersion", ip))
		// if err != nil || status != 200 {
		// 	return c, fmt.Errorf("status: %d, error: %w", status, err)
		// }

		// s = strings.Split(body, "=")
		// s = strings.Split(s[1], ",")
		// c.Software = s[0]

		return c, nil

	case 404:
		// Parse Modelo --------------------------------
		_, status, err := Requisitador(t, fmt.Sprintf("http://%s/axis-cgi/param.cgi?action=list&group=Brand.ProdShortName", ip))
		if err != nil || status != 200 {
			return c, fmt.Errorf("status: %d, error: %w", status, err)
		}

		switch status {
		case 401:
			return c, errors.New("erro na autenticação - verificar usuário e senha")

		case 200:
			c.Manufacturer = "axis"
			response = getURLs("axis")

			// // Parse Modelo --------------------------------
			// s := strings.Split(body, "=")
			// c.Model = s[1]

			// // Parse Serial Number -------------------------
			// body, status, err = Requisitador(t, fmt.Sprintf("http://%s/axis-cgi/param.cgi?action=list&group=Properties.System.SerialNumber", ip))
			// if err != nil || status != 200 {
			// 	return c, fmt.Errorf("status: %d, error: %w", status, err)
			// }
			// s = strings.Split(body, "=")
			// c.serialNumber = strings.ToUpper(s[1])

			// // Parse MAC Number ----------------------------
			// body, status, err = Requisitador(t, fmt.Sprintf("http://%s/axis-cgi/param.cgi?action=list&group=Network.eth0.MACAddress", ip))
			// if err != nil || status != 200 {
			// 	return c, fmt.Errorf("status: %d, error: %w", status, err)
			// }
			// s = strings.Split(body, "=")
			// c.MAC = strings.ReplaceAll(strings.ToUpper(s[1]), ":", "")

			// // Parse Firmware Number -----------------------
			// body, status, err = Requisitador(t, fmt.Sprintf("http://%s/axis-cgi/param.cgi?action=list&group=Properties.Firmware.Version", ip))
			// if err != nil || status != 200 {
			// 	return c, fmt.Errorf("status: %d, error: %w", status, err)
			// }
			// s = strings.Split(body, "=")
			// c.Software = strings.ToUpper(s[1])

		case 404:
			// TODO testar novo VAPIX (tem que utilizar POST Method)
			return c, errors.New("dispositivo desconhecido")

		default:
			return c, fmt.Errorf("status: %d, error: status desconhecido", status)

		}

	default:
		return c, fmt.Errorf("status: %d, error: status desconhecido", status)

	}

	for k, v := range response {
		body, status, err := Requisitador(t, fmt.Sprintf("http:/%s%s", ip, v))
		if err != nil || status != 200 {
			return c, fmt.Errorf("status: %d, error: %w", status, err)
		}
		s := strings.Split(body, "=")
		result := strings.ReplaceAll(strings.ToUpper(s[1]), ":", "")

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
