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

// const (
// 	condition = iota
// 	video
// 	title
// 	network
// 	date
// 	userAdd
// 	userAddSigeo
// 	changePass
// )

// TODO exposure e outros parametros não estão funcionando

func NewDahuaConfigurator() map[string]string {
	return map[string]string{
		// condition:    "/cgi-bin/configManager.cgi?action=setConfig&VideoInBacklight[0][2].Mode=WideDynamic&VideoInBacklight[0][2].WideDynamicRange=20&VideoInExposure[0][2].Compensation=20&VideoInDayNight[0][2].Mode=Color&VideoInMode[0].Config[0]=2",
		// video:        "/cgi-bin/configManager.cgi?action=setConfig&Encode[0].MainFormat[0].Video.Compression=H.264&Encode[0].MainFormat[0].Video.Profile=Main&Encode[0].MainFormat[0].Video.resolution=1920x1080&Encode[0].MainFormat[0].Video.Height=1080&Encode[0].MainFormat[0].Video.Width=1920&Encode[0].MainFormat[0].Video.FPS=16&Encode[0].MainFormat[0].Video.BitRateControl=CBR&Encode[0].MainFormat[0].Video.BitRate=4096&Encode[0].MainFormat[0].Video.GOP=32",

		"title": "/cgi-bin/configManager.cgi?action=setConfig&ChannelTitle[0].Name=SSP%2FDF%20-%20%localizacao%%20-%20%ponto%&table.VideoWidget[0].FontSize=32",

		"channelTitle":      "/cgi-bin/configManager.cgi?action=setConfig&table.VideoWidget[0].ChannelTitle.FrontColor[0]=255&table.VideoWidget[0].ChannelTitle.FrontColor[1]=255&table.VideoWidget[0].ChannelTitle.FrontColor[2]=0&table.VideoWidget[0].ChannelTitle.FrontColor[3]=0",
		"covers0":           "/cgi-bin/configManager.cgi?action=setConfig&table.VideoWidget[0].Covers[0].FrontColor[0]=255&table.VideoWidget[0].Covers[0].FrontColor[1]=255&table.VideoWidget[0].Covers[0].FrontColor[2]=0&table.VideoWidget[0].Covers[0].FrontColor[3]=0",
		"covers1":           "/cgi-bin/configManager.cgi?action=setConfig&table.VideoWidget[0].Covers[1].FrontColor[0]=255&table.VideoWidget[0].Covers[1].FrontColor[1]=255&table.VideoWidget[0].Covers[1].FrontColor[2]=0&table.VideoWidget[0].Covers[1].FrontColor[3]=0",
		"covers2":           "/cgi-bin/configManager.cgi?action=setConfig&table.VideoWidget[0].Covers[2].FrontColor[0]=255&table.VideoWidget[0].Covers[2].FrontColor[1]=255&table.VideoWidget[0].Covers[2].FrontColor[2]=0&table.VideoWidget[0].Covers[2].FrontColor[3]=0",
		"covers3":           "/cgi-bin/configManager.cgi?action=setConfig&table.VideoWidget[0].Covers[3].FrontColor[0]=255&table.VideoWidget[0].Covers[3].FrontColor[1]=255&table.VideoWidget[0].Covers[3].FrontColor[2]=0&table.VideoWidget[0].Covers[3].FrontColor[3]=0",
		"covers4":           "/cgi-bin/configManager.cgi?action=setConfig&table.VideoWidget[0].Covers[4].FrontColor[0]=255&table.VideoWidget[0].Covers[4].FrontColor[1]=255&table.VideoWidget[0].Covers[4].FrontColor[2]=0&table.VideoWidget[0].Covers[4].FrontColor[3]=0",
		"covers5":           "/cgi-bin/configManager.cgi?action=setConfig&table.VideoWidget[0].Covers[5].FrontColor[0]=255&table.VideoWidget[0].Covers[5].FrontColor[1]=255&table.VideoWidget[0].Covers[5].FrontColor[2]=0&table.VideoWidget[0].Covers[5].FrontColor[3]=0",
		"covers6":           "/cgi-bin/configManager.cgi?action=setConfig&table.VideoWidget[0].Covers[6].FrontColor[0]=255&table.VideoWidget[0].Covers[6].FrontColor[1]=255&table.VideoWidget[0].Covers[6].FrontColor[2]=0&table.VideoWidget[0].Covers[6].FrontColor[3]=0",
		"covers7":           "/cgi-bin/configManager.cgi?action=setConfig&table.VideoWidget[0].Covers[7].FrontColor[0]=255&table.VideoWidget[0].Covers[7].FrontColor[1]=255&table.VideoWidget[0].Covers[7].FrontColor[2]=0&table.VideoWidget[0].Covers[7].FrontColor[3]=0",
		"customTitle0":      "/cgi-bin/configManager.cgi?action=setConfig&table.VideoWidget[0].CustomTitle[0].FrontColor[0]=255&table.VideoWidget[0].CustomTitle[0].FrontColor[1]=255&table.VideoWidget[0].CustomTitle[0].FrontColor[2]=0&table.VideoWidget[0].CustomTitle[0].FrontColor[3]=0",
		"customTitle1":      "/cgi-bin/configManager.cgi?action=setConfig&table.VideoWidget[0].CustomTitle[1].FrontColor[0]=255&table.VideoWidget[0].CustomTitle[1].FrontColor[1]=255&table.VideoWidget[0].CustomTitle[1].FrontColor[2]=0&table.VideoWidget[0].CustomTitle[1].FrontColor[3]=0",
		"customTitle2":      "/cgi-bin/configManager.cgi?action=setConfig&table.VideoWidget[0].CustomTitle[2].FrontColor[0]=255&table.VideoWidget[0].CustomTitle[2].FrontColor[1]=255&table.VideoWidget[0].CustomTitle[2].FrontColor[2]=0&table.VideoWidget[0].CustomTitle[2].FrontColor[3]=0",
		"customTitle3":      "/cgi-bin/configManager.cgi?action=setConfig&table.VideoWidget[0].CustomTitle[3].FrontColor[0]=255&table.VideoWidget[0].CustomTitle[3].FrontColor[1]=255&table.VideoWidget[0].CustomTitle[3].FrontColor[2]=0&table.VideoWidget[0].CustomTitle[3].FrontColor[3]=0",
		"osdMobileState":    "/cgi-bin/configManager.cgi?action=setConfig&table.VideoWidget[0].OSDMobileState.FrontColor[0]=255&table.VideoWidget[0].OSDMobileState.FrontColor[1]=255&table.VideoWidget[0].OSDMobileState.FrontColor[2]=0&table.VideoWidget[0].OSDMobileState.FrontColor[3]=0",
		"ptzCoordinates":    "/cgi-bin/configManager.cgi?action=setConfig&table.VideoWidget[0].PTZCoordinates.FrontColor[0]=255&table.VideoWidget[0].PTZCoordinates.FrontColor[1]=255&table.VideoWidget[0].PTZCoordinates.FrontColor[2]=0&table.VideoWidget[0].PTZCoordinates.FrontColor[3]=0",
		"ptzDirection":      "/cgi-bin/configManager.cgi?action=setConfig&table.VideoWidget[0].PTZDirection.FrontColor[0]=255&table.VideoWidget[0].PTZDirection.FrontColor[1]=255&table.VideoWidget[0].PTZDirection.FrontColor[2]=0&table.VideoWidget[0].PTZDirection.FrontColor[3]=0",
		"ptzOSDMenu":        "/cgi-bin/configManager.cgi?action=setConfig&table.VideoWidget[0].PTZOSDMenu.FrontColor[0]=255&table.VideoWidget[0].PTZOSDMenu.FrontColor[1]=255&table.VideoWidget[0].PTZOSDMenu.FrontColor[2]=0&table.VideoWidget[0].PTZOSDMenu.FrontColor[3]=0",
		"ptzOSDMenuViaApp":  "/cgi-bin/configManager.cgi?action=setConfig&table.VideoWidget[0].PTZOSDMenuViaApp.FrontColor[0]=255&table.VideoWidget[0].PTZOSDMenuViaApp.FrontColor[1]=255&table.VideoWidget[0].PTZOSDMenuViaApp.FrontColor[2]=0&table.VideoWidget[0].PTZOSDMenuViaApp.FrontColor[3]=0",
		"ptzPreset":         "/cgi-bin/configManager.cgi?action=setConfig&table.VideoWidget[0].PTZPreset.FrontColor[0]=255&table.VideoWidget[0].PTZPreset.FrontColor[1]=255&table.VideoWidget[0].PTZPreset.FrontColor[2]=0&table.VideoWidget[0].PTZPreset.FrontColor[3]=3",
		"ptzZoom":           "/cgi-bin/configManager.cgi?action=setConfig&table.VideoWidget[0].PTZZoom.FrontColor[0]=255&table.VideoWidget[0].PTZZoom.FrontColor[1]=255&table.VideoWidget[0].PTZZoom.FrontColor[2]=0&table.VideoWidget[0].PTZZoom.FrontColor[3]=0",
		"pictureTitle":      "/cgi-bin/configManager.cgi?action=setConfig&table.VideoWidget[0].PictureTitle.FrontColor[0]=255&table.VideoWidget[0].PictureTitle.FrontColor[1]=255&table.VideoWidget[0].PictureTitle.FrontColor[2]=0&table.VideoWidget[0].PictureTitle.FrontColor[3]=0",
		"ptzPattern":        "/cgi-bin/configManager.cgi?action=setConfig&table.VideoWidget[0].PtzPattern.FrontColor[0]=255&table.VideoWidget[0].PtzPattern.FrontColor[1]=255&table.VideoWidget[0].PtzPattern.FrontColor[2]=0&table.VideoWidget[0].PtzPattern.FrontColor[3]=0",
		"ptzRS485Detect":    "/cgi-bin/configManager.cgi?action=setConfig&table.VideoWidget[0].PtzRS485Detect.FrontColor[0]=255&table.VideoWidget[0].PtzRS485Detect.FrontColor[1]=255&table.VideoWidget[0].PtzRS485Detect.FrontColor[2]=0&table.VideoWidget[0].PtzRS485Detect.FrontColor[3]=0",
		"temperature":       "/cgi-bin/configManager.cgi?action=setConfig&table.VideoWidget[0].Temperature.FrontColor[0]=255&table.VideoWidget[0].Temperature.FrontColor[1]=255&table.VideoWidget[0].Temperature.FrontColor[2]=0&table.VideoWidget[0].Temperature.FrontColor[3]=0",
		"timeTitle":         "/cgi-bin/configManager.cgi?action=setConfig&table.VideoWidget[0].TimeTitle.FrontColor[0]=255&table.VideoWidget[0].TimeTitle.FrontColor[1]=255&table.VideoWidget[0].TimeTitle.FrontColor[2]=0&table.VideoWidget[0].TimeTitle.FrontColor[3]=0",
		"trafficFlowTitle":  "/cgi-bin/configManager.cgi?action=setConfig&table.VideoWidget[0].TrafficFlowTitle.FrontColor[0]=255&table.VideoWidget[0].TrafficFlowTitle.FrontColor[1]=255&table.VideoWidget[0].TrafficFlowTitle.FrontColor[2]=0&table.VideoWidget[0].TrafficFlowTitle.FrontColor[3]=0",
		"userDefinedTitle0": "/cgi-bin/configManager.cgi?action=setConfig&table.VideoWidget[0].UserDefinedTitle[0].FrontColor[0]=255&table.VideoWidget[0].UserDefinedTitle[0].FrontColor[1]=255&table.VideoWidget[0].UserDefinedTitle[0].FrontColor[2]=0&table.VideoWidget[0].UserDefinedTitle[0].FrontColor[3]=0",
		"userDefinedTitle1": "/cgi-bin/configManager.cgi?action=setConfig&table.VideoWidget[0].UserDefinedTitle[1].FrontColor[0]=255&table.VideoWidget[0].UserDefinedTitle[1].FrontColor[1]=255&table.VideoWidget[0].UserDefinedTitle[1].FrontColor[2]=0&table.VideoWidget[0].UserDefinedTitle[1].FrontColor[3]=0",
		"userDefinedTitle2": "/cgi-bin/configManager.cgi?action=setConfig&table.VideoWidget[0].UserDefinedTitle[2].FrontColor[0]=255&table.VideoWidget[0].UserDefinedTitle[2].FrontColor[1]=255&table.VideoWidget[0].UserDefinedTitle[2].FrontColor[2]=0&table.VideoWidget[0].UserDefinedTitle[2].FrontColor[3]=0",
		"userDefinedTitle3": "/cgi-bin/configManager.cgi?action=setConfig&table.VideoWidget[0].UserDefinedTitle[3].FrontColor[0]=255&table.VideoWidget[0].UserDefinedTitle[3].FrontColor[1]=255&table.VideoWidget[0].UserDefinedTitle[3].FrontColor[2]=0&table.VideoWidget[0].UserDefinedTitle[3].FrontColor[3]=0",
		"voltageStatus":     "/cgi-bin/configManager.cgi?action=setConfig&table.VideoWidget[0].VoltageStatus.FrontColor[0]=255&table.VideoWidget[0].VoltageStatus.FrontColor[1]=255&table.VideoWidget[0].VoltageStatus.FrontColor[2]=0&table.VideoWidget[0].VoltageStatus.FrontColor[3]=0",

		// date:         "/cgi-bin/configManager.cgi?action=setConfig&NTP.TimeZone=22&NTP.TimeZoneDesc=Brasilia&NTP.Address=10.92.0.54&NTP.Enable=true",
		//userAdd:      "/cgi-bin/userManager.cgi?action=addUser&user.Name=%userAddAdmin%&user.Password=%passUserAddAdmin%A&user.Group=admin&user.Sharable=true&user.Reserved=false",
		//userAddSigeo: "/cgi-bin/userManager.cgi?action=addUser&user.Name=%userAddUser%&user.Password=%passUserAddUser%A&user.Group=user&user.Sharable=true&user.Reserved=false",
		//changePass:   "/cgi-bin/userManager.cgi?action=modifyPassword&name=admin&pwd=%senhaMaster%&pwdOld=%senhaAntiga%",
		//network:      "/cgi-bin/configManager.cgi?action=setConfig&Network.Hostname=%patrimonio%&Network.eth0.IPAddress=%ip%&Network.eth0.DefaultGateway=%gateway%&Network.eth0.SubnetMask=%mascara%",
	}
}
