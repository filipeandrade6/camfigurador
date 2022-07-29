package main

import tea "github.com/charmbracelet/bubbletea"

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.stage {
	case credentials:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "ctrl+c", "esc":
				return m, tea.Quit

			case "tab", "shift+tab", "enter", "up", "down":
				s := msg.String()

				// Pressed OK button save to model and goes to configuration config
				if s == "enter" && m.focusIndexCredentials == len(m.inputsCredentials) {
					m.credentialsToConfiguration()
					m.getConfiguration()
					return m, nil
				}

				if s == "up" || s == "shift+tab" {
					m.focusIndexCredentials--
				} else {
					m.focusIndexCredentials++
				}

				if m.focusIndexCredentials > len(m.inputsCredentials) {
					m.focusIndexCredentials = 0
				} else if m.focusIndexCredentials < 0 {
					m.focusIndexCredentials = len(m.inputsCredentials)
				}

				cmds := make([]tea.Cmd, len(m.inputsCredentials))
				for i := 0; i <= len(m.inputsCredentials)-1; i++ {
					if i == m.focusIndexCredentials {
						cmds[i] = m.inputsCredentials[i].Focus()
						m.inputsCredentials[i].PromptStyle = focusedStyle
						m.inputsCredentials[i].TextStyle = focusedStyle
						continue
					}
					// remove focused state
					m.inputsCredentials[i].Blur()
					m.inputsCredentials[i].PromptStyle = noStyle
					m.inputsCredentials[i].TextStyle = noStyle
				}

				return m, tea.Batch(cmds...)
			}
		}

		// Handle character input and blinking
		cmd := m.updateInputsCredentials(msg)

		return m, cmd

	case configuration:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "ctrl+c", "esc":
				m.configurationToCredentials()
				return m, nil

			case "tab", "shift+tab", "enter", "up", "down":
				s := msg.String()

				// Pressed OK button save to model and goes to configuration config
				if s == "enter" && m.focusIndexConfiguration == len(m.inputsConfiguration) {
					m.getConfiguration()
					return m, nil
				}

				if s == "up" || s == "shift+tab" {
					m.focusIndexConfiguration--
				} else {
					m.focusIndexConfiguration++
				}

				if m.focusIndexConfiguration > len(m.inputsConfiguration) {
					m.focusIndexConfiguration = 0
				} else if m.focusIndexConfiguration < 0 {
					m.focusIndexConfiguration = len(m.inputsConfiguration)
				}

				cmds := make([]tea.Cmd, len(m.inputsConfiguration))
				for i := 0; i <= len(m.inputsConfiguration)-1; i++ {
					if i == m.focusIndexConfiguration {
						cmds[i] = m.inputsConfiguration[i].Focus()
						m.inputsConfiguration[i].PromptStyle = focusedStyle
						m.inputsConfiguration[i].TextStyle = focusedStyle
						continue
					}
					// remove focused state
					m.inputsConfiguration[i].Blur()
					m.inputsConfiguration[i].PromptStyle = noStyle
					m.inputsConfiguration[i].TextStyle = noStyle
				}

				return m, tea.Batch(cmds...)
			}
		}

		// Handle character input and blinking
		cmd := m.updateInputsConfiguration(msg)

		return m, cmd
	}

	return m, nil
}
