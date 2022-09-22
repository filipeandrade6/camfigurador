package main

import "strings"

// type ConfigurationCfg struct {
// 	IPAddr       string
// 	Gateway      string
// 	SubnetMask   string
// 	ChannelTitle string
// 	Hostname     string
// 	Ponto        string
// }

type URLer struct {
	urlBase string
	items   []string
}

func (u *URLer) interpolar(ip string) string {
	return "http://" + ip + u.urlBase
}

func (m *model) Configurar() error {
	cfg := make(map[int]string)

	switch m.response.Manufacturer {
	case "axis":

	case "dahua":
		cfg = NewDahuaConfigurator()

	default:
	}

	for k, v := range cfg {

		url := "http://" + m.deviceIP + strings.Replace(v, '%'+k+'%', k)

		body, statusCode, err := m.Requisitador(url)
		if err != nil {
			return c, err
		}
	}

	return nil
}
