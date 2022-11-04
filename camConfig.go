package main

import (
	"fmt"
	"net/url"
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
		u := "http://" + m.deviceIP + v

		// video
		l := &url.URL{Path: m.inputsConfiguration[3].Value()}
		p := &url.URL{Path: m.inputsConfiguration[4].Value()}
		u = strings.ReplaceAll(u, "%localizacao%", l.String())
		u = strings.ReplaceAll(u, "%ponto%", p.String())

		// userAdd
		u = strings.ReplaceAll(u, "%userAddAdmin%", "outro")
		u = strings.ReplaceAll(u, "%passUserAddAdmin%", "8t%25fbNRFpS80lGqo")

		// userAddSigeo
		u = strings.ReplaceAll(u, "%userAddUser%", "outro2")
		u = strings.ReplaceAll(u, "%passUserAddUser%", "8t%25fbNRFpS80lGqo")

		// changePass
		u = strings.ReplaceAll(u, "%senhaMaster%", "abc")
		u = strings.ReplaceAll(u, "%senhaAntiga%", "def")

		// network
		// TODO trocar o IP vai ser necessário alterar o IP base
		// TODO criar novo campo no model para novo ip?
		u = strings.ReplaceAll(u, "%patrimonio%", url.QueryEscape(m.inputsConfiguration[5].Value()))
		u = strings.ReplaceAll(u, "%ip%", url.QueryEscape(m.inputsConfiguration[0].Value()))
		u = strings.ReplaceAll(u, "%gateway%", url.QueryEscape(m.inputsConfiguration[1].Value()))
		u = strings.ReplaceAll(u, "%mascara%", url.QueryEscape(m.inputsConfiguration[2].Value()))

		fmt.Println(u)

		body, statusCode, err := m.Requisitador(u)
		if err != nil {
			return err
		}

		if statusCode != 200 {
			return fmt.Errorf("não é status 200 - code: %d - body: %s", statusCode, body)
		}
	}

	return nil
}
