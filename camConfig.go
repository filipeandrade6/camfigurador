package main

import (
	"fmt"
	"strings"
)

// TODO codificar tudo para URL
// TODO preparar para AXIS
// TODO criar API separada para integrar com Bot Telegram ou Mattermost

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
		url = strings.ReplaceAll(url, "%passUserAddAdmin%", "8t%25fbNRFpS80lGqo")

		// userAddSigeo
		url = strings.ReplaceAll(url, "%userAddUser%", "outro2")
		url = strings.ReplaceAll(url, "%passUserAddUser%", "8t%25fbNRFpS80lGqo")

		// changePass
		url = strings.ReplaceAll(url, "%senhaMaster%", "abc")
		url = strings.ReplaceAll(url, "%senhaAntiga%", "def")

		// network
		// TODO trocar o IP vai ser necessário alterar o IP base
		// TODO criar novo campo no model para novo ip?
		url = strings.ReplaceAll(url, "%patrimonio%", m.inputsConfiguration[5].Value())
		url = strings.ReplaceAll(url, "%ip%", m.inputsConfiguration[0].Value())
		url = strings.ReplaceAll(url, "%gateway%", m.inputsConfiguration[1].Value())
		url = strings.ReplaceAll(url, "%mascara%", m.inputsConfiguration[2].Value())

		fmt.Println(url)

		body, statusCode, err := m.Requisitador(url)
		if err != nil {
			return err
		}

		if statusCode != 200 {
			return fmt.Errorf("não é status 200 - code: %d - body: %s", statusCode, body)
		}
	}

	return nil
}
