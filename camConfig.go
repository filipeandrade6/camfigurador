package main

import (
	"errors"
	"strings"
)

// type ConfigurationCfg struct {
// 	IPAddr       string
// 	Gateway      string
// 	SubnetMask   string
// 	ChannelTitle string
// 	Hostname     string
// 	Ponto        string
// }

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
		url = strings.ReplaceAll(url, "%userAddAdmin%", "ditec")
		url = strings.ReplaceAll(url, "%passUserAddAdmin%", "DITECam%23%7B8863%7D")

		// userAddSigeo
		url = strings.ReplaceAll(url, "%userAddUser%", "sigeo")
		url = strings.ReplaceAll(url, "%passUserAddUser%", "sigeo%402018")

		// chagePass
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
