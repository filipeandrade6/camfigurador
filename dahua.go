package main

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
	title
	network
	date
	userAdd
	userAddSigeo
	changePass
)

// TODO exposure e outros parametros não estão funcionando

func NewDahuaConfigurator() map[int]string {
	return map[int]string{
		condition:    "/cgi-bin/configManager.cgi?action=setConfig&VideoInBacklight[0][2].Mode=WideDynamic&VideoInBacklight[0][2].WideDynamicRange=20&VideoInExposure[0][2].Compensation=20&VideoInDayNight[0][2].Mode=Color&VideoInMode[0].Config[0]=2",
		video:        "/cgi-bin/configManager.cgi?action=setConfig&Encode[0].MainFormat[0].Video.Compression=H.264&Encode[0].MainFormat[0].Video.Profile=Main&Encode[0].MainFormat[0].Video.resolution=1920x1080&Encode[0].MainFormat[0].Video.Height=1080&Encode[0].MainFormat[0].Video.Width=1920&Encode[0].MainFormat[0].Video.FPS=16&Encode[0].MainFormat[0].Video.BitRateControl=CBR&Encode[0].MainFormat[0].Video.BitRate=4096&Encode[0].MainFormat[0].Video.GOP=32",
		title:        "/cgi-bin/configManager.cgi?action=setConfig&ChannelTitle[0].Name=SSP%5CDF%20-%20%localizacao%%20-%20%ponto%&table.VideoWidget[0].FontSize=32&table.VideoWidget[0].ChannelTitle.FrontColor[0]=255&table.VideoWidget[0].ChannelTitle.FrontColor[1]=255&table.VideoWidget[0].ChannelTitle.FrontColor[2]=0&table.VideoWidget[0].ChannelTitle.FrontColor[3]=0",
		date:         "/cgi-bin/configManager.cgi?action=setConfig&NTP.TimeZone=22&NTP.TimeZoneDesc=Brasilia&NTP.Address=10.92.0.54&NTP.Enable=true",
		userAdd:      "/cgi-bin/userManager.cgi?action=addUser&user.Name=%userAddAdmin%&user.Password=%passUserAddAdmin%A&user.Group=admin&user.Sharable=true&user.Reserved=false",
		userAddSigeo: "/cgi-bin/userManager.cgi?action=addUser&user.Name=%userAddUser%&user.Password=%passUserAddUser%A&user.Group=user&user.Sharable=true&user.Reserved=false",
		// changePass:   "/cgi-bin/userManager.cgi?action=modifyPassword&name=admin&pwd=%senhaMaster%&pwdOld=%senhaAntiga%",
		// network: "/cgi-bin/configManager.cgi?action=setConfig&Network.Hostname=%patrimonio%&Network.eth0.IPAddress=%ip%&Network.eth0.DefaultGateway=%gateway%&Network.eth0.SubnetMask=%mascara%",
	}
}
