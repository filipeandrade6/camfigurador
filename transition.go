package main

func (m *model) credentialsToConfiguration() {
	m.deviceIP = m.inputsCredentials[0].Value() // TODO trocar o index de número para constante
	m.httpTransport.Username = m.inputsCredentials[1].Value()
	m.httpTransport.Password = m.inputsCredentials[2].Value()

	camCfg, err := m.IdentificadorDeModelo()
	if err != nil {
		m.err = err
		return
	}

	// TODO colocar um loading até terinar a requisição?

	m.response = camCfg

	m.stage = configuration
	m.focusIndexCredentials = 0
	m.inputsConfiguration[0].Focus()
	m.inputsConfiguration[0].PromptStyle = focusedStyle
	m.inputsConfiguration[0].TextStyle = focusedStyle
}

func (m *model) configurationToCredentials() {
	m.stage = credentials // TODO TROCAR
	m.focusIndexConfiguration = 0
	m.inputsCredentials[0].Focus()
	m.inputsCredentials[0].PromptStyle = focusedStyle
	m.inputsCredentials[0].TextStyle = focusedStyle
}
