package main

import (
	"net/http"
	"time"

	dac "github.com/xinsnake/go-http-digest-auth-client"
)

var baseUrl = "/cgi-bin/configManager.cgi?action=setConfig&"

type Configurator struct {
	Condition string
	Video     string
	Network   string
	PTZ       string
	Date      string
	Account   string
}

func NewDahuaConfigurator() Configurator {
	return Configurator{
		Condition: "VideoInBacklight[0][2].Mode=WideDynamic&VideoInBacklight[0][2].WideDynamicRange=20&VideoInExposure[0][2].Compensation=20&VideoInDayNight[0][2].Mode=Color&VideoInMode[0].Config[0]=2",

		Video:   "Encode[0].MainFormat[0].Video.Compression=H.264&Encode[0].MainFormat[0].Video.Profile=Main&Encode[0].MainFormat[0].Video.resolution=1920x1080&Encode[0].MainFormat[0].Video.Height=1080&Encode[0].MainFormat[0].Video.Width=1920&Encode[0].MainFormat[0].Video.Width=1920&Encode[0].MainFormat[0].Video.FPS=16&Encode[0].MainFormat[0].Video.BitRateControl=CBR&Encode[0].MainFormat[0].Video.BitRate=4096&Encode[0].MainFormat[0].Video.GOP=32&Encode[0].MainFormat[1].VideoEnable=False&ChannelTitle[0].Name={SSP/DF - localização - PONTO}",
		Network: "Network.Hostname={nome da câmera}&Network.eth0.IPAddress={ip}&Network.eth0.DefaultGateway={gateway}&Network.eth0.SubnetMask={mascara de subrede}",
		Date:    "NTP.TimeZone=22&NTP.TimeZoneDesc=Brasilia&NTP.Address=10.92.0.54&NTP.Enable=true", // OK
		// 	Account:
		// 		"url base":                           "http://<ip-address>/cgi-bin/userManager.cgi?action=", // TODO
		// 		"add account":                        "addUser&user.Name={usuario}&us er.Password={senha}&user.Group={user|admin}&user.Sharable=true&user.Reserved=false",
		// 		"change admin password":              "modifyPassword&name=admin&pwd={senha}&pwdOld={senha-antiga}",
		// 		"onvif user / add account":           "<<ACHO QUE NÃO TEM>>", // TODO
		// 		"onvif user / change admin password": "modifyPasswordByManager&userName=admin&pwd{senha}&managerName={admin}&managerPwd={senha}&accountType=1",
		// }
	}
}

func configurar(ip, user, password string, cfg Configurator) error {
	t := dac.NewTransport(user, password)
	t.HTTPClient = &http.Client{Timeout: 2 * time.Second}

	_, statusCode, err := Requisitador(t, url)
	if err != nil {
		return c, err
	}

	return nil
}
