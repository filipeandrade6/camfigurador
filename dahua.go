package main

import (
	"net/http"
	"time"

	dac "github.com/xinsnake/go-http-digest-auth-client"
)

var baseUrl = "/cgi-bin/configManager.cgi?action=setConfig&"

type Configurator struct {
	Condition    string
	Video        string
	Network      string
	PTZ          string
	Date         string
	Account      string
	UserAdd      string
	UserAddSigeo string
	ChangePass   string
}

const (
	condition = iota
	video
	network
	date
	userAdd
	userAddSigeo
	changePass
)

func NewDahuaConfigurator() map[int]string {
	return map[int]string{
		condition:    "/cgi-bin/configManager.cgi?action=setConfig&VideoInBacklight[0][2].Mode=WideDynamic&VideoInBacklight[0][2].WideDynamicRange=20&VideoInExposure[0][2].Compensation=20&VideoInDayNight[0][2].Mode=Color&VideoInMode[0].Config[0]=2",
		video:        "/cgi-bin/configManager.cgi?action=setConfig&Encode[0].MainFormat[0].Video.Compression=H.264&Encode[0].MainFormat[0].Video.Profile=Main&Encode[0].MainFormat[0].Video.resolution=1920x1080&Encode[0].MainFormat[0].Video.Height=1080&Encode[0].MainFormat[0].Video.Width=1920&Encode[0].MainFormat[0].Video.FPS=16&Encode[0].MainFormat[0].Video.BitRateControl=CBR&Encode[0].MainFormat[0].Video.BitRate=4096&Encode[0].MainFormat[0].Video.GOP=32&ChannelTitle[0].Name=%localizacao%&table.VideoWidget[0].FontSize=32&table.VideoWidget[0].ChannelTitle.FrontColor[0]=255&table.VideoWidget[0].ChannelTitle.FrontColor[1]=255&table.VideoWidget[0].ChannelTitle.FrontColor[2]=0&table.VideoWidget[0].ChannelTitle.FrontColor[3]=0",
		network:      "/cgi-bin/configManager.cgi?action=setConfig&Network.Hostname=%patrimonio%&Network.eth0.IPAddress=%ip%&Network.eth0.DefaultGateway=%gateway%&Network.eth0.SubnetMask=%mascara%",
		date:         "/cgi-bin/configManager.cgi?action=setConfig&NTP.TimeZone=22&NTP.TimeZoneDesc=Brasilia&NTP.Address=10.92.0.54&NTP.Enable=true",
		userAdd:      "/cgi-bin/userManager.cgi?action=addUser&user.Name=%userAddAdmin%&user.Password=%passUserAddAdmin%A&user.Group=admin&user.Sharable=true&user.Reserved=false",
		userAddSigeo: "/cgi-bin/userManager.cgi?action=addUser&user.Name=%userAddUser%&user.Password=%passUserAddUser%A&user.Group=user&user.Sharable=true&user.Reserved=false",
		changePass:   "/cgi-bin/userManager.cgi?action=modifyPassword&name=admin&pwd=%senhaMaster%&pwdOld=%senhaAntiga%",
	}
}

func (m model) configurar() error {
	t := dac.NewTransport(m.user, m.pass)
	t.HTTPClient = &http.Client{Timeout: 2 * time.Second}

	cfg := make(map[int]string)

	switch m.response.Manufacturer {
	case "axis":

	case "dahua":
		cfg = NewDahuaConfigurator()

	default:
	}

	for k, v := range cfg {
		switch v {
			case
		}
		body, statusCode, err := Requisitador(t, url)
		if err != nil {
			return c, err
		}
	}

	return nil
}
