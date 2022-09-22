package main

import tea "github.com/charmbracelet/bubbletea"

func (m *model) updateInputsCredentials(msg tea.Msg) tea.Cmd {
	var cmds = make([]tea.Cmd, len(m.inputsCredentials))

	// Only text inputsCredentials with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputsCredentials {
		m.inputsCredentials[i], cmds[i] = m.inputsCredentials[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m *model) updateInputsConfiguration(msg tea.Msg) tea.Cmd {
	var cmds = make([]tea.Cmd, len(m.inputsConfiguration))

	// Only text inputsConfiguration with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputsConfiguration {
		m.inputsConfiguration[i], cmds[i] = m.inputsConfiguration[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

// func (m *model) getConfiguration() {
// 	camCfg, err := IdentificadorDeModelo(
// 		m.inputsCredentials[0].Value(),
// 		m.inputsCredentials[1].Value(),
// 		m.inputsCredentials[2].Value(),
// 	)
// 	if err != nil {
// 		m.err = err
// 		return
// 	}

// 	m.response = camCfg
// }
