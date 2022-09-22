package main

func (m *model) credentialsToConfiguration() {
	camCfg, err := IdentificadorDeModelo(
		m.inputsCredentials[0].Value(),
		m.inputsCredentials[1].Value(),
		m.inputsCredentials[2].Value(),
	)
	if err != nil {
		m.err = err
		return
	}

	m.response = camCfg

	m.stage = configuration
	m.focusIndexCredentials = 0
	m.inputsConfiguration[0].Focus()
	m.inputsConfiguration[0].PromptStyle = focusedStyle
	m.inputsConfiguration[0].TextStyle = focusedStyle
}

func (m *model) configurationToCredentials() {
	m.stage = credentials
	m.focusIndexConfiguration = 0
	m.inputsCredentials[0].Focus()
	m.inputsCredentials[0].PromptStyle = focusedStyle
	m.inputsCredentials[0].TextStyle = focusedStyle
}
