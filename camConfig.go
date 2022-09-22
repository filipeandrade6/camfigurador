package main

import (
	"errors"
	"strings"
)

func (m *model) Configurar() error {
	cfg := make(map[int]string)

	switch m.response.Manufacturer {
	case "axis":

	case "dahua":
		cfg = NewDahuaConfigurator()

	default:
	}

	for _, v := range cfg {
		url := "http://" + m.deviceIP + v

		// video
		url = strings.ReplaceAll(url, "%localizacao%", m.inputsConfiguration[3].Value())
		url = strings.ReplaceAll(url, "%ponto%", m.inputsConfiguration[4].Value())

		// userAdd
		url = strings.ReplaceAll(url, "%userAddAdmin%", "outro")
		url = strings.ReplaceAll(url, "%passUserAddAdmin%", "outro")

		// userAddSigeo
		url = strings.ReplaceAll(url, "%userAddUser%", "outro2")
		url = strings.ReplaceAll(url, "%passUserAddUser%", "outro2")

		// changePass
		url = strings.ReplaceAll(url, "%senhaMaster%", "abc")
		url = strings.ReplaceAll(url, "%senhaAntiga%", "def")

		// network
		// TODO trocar o IP vai ser necessário alterar o IP base
		url = strings.ReplaceAll(url, "%patrimonio%", m.inputsConfiguration[5].Value())
		url = strings.ReplaceAll(url, "%ip%", m.inputsConfiguration[0].Value())
		url = strings.ReplaceAll(url, "%gateway%", m.inputsConfiguration[1].Value())
		url = strings.ReplaceAll(url, "%mascara%", m.inputsConfiguration[2].Value())

		_, statusCode, err := m.Requisitador(url)
		if err != nil {
			return err
		}

		if statusCode != 200 {
			return errors.New("não é status 200")
		}
	}

	return nil
}
