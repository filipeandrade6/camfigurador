package main

import (
	"fmt"
	"strings"
)

func (m model) View() string {
	var b strings.Builder

	b.WriteString("\n CONFIGURADOR - DE - CÂMERA\n")
	b.WriteString("Filipe Andrade - filipe.engenharia42@gmail.com\n\n")

	switch m.stage {
	case credentials:
		b.WriteString("[Ctrl+C] ou ESC para sair\n\n")

		if m.err != nil {
			fmt.Fprintf(&b, ">>> %s\n\n", m.err)
		}

		for i := range m.inputsCredentials {
			b.WriteString(m.inputsCredentials[i].View())
			if i < len(m.inputsCredentials)-1 {
				b.WriteRune('\n')
			}
		}

		button := &blurredButton
		if m.focusIndexCredentials == len(m.inputsCredentials) {
			button = &focusedButton
		}
		fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	case configuration:
		if m.err != nil {
			return b.String() // TODO arrumar isso aqui
		}

		fmt.Fprintf(&b, "      fabricante: %s\n          modelo: %s\n             MAC: %s\n número de série: %s\n        software: %s\n\n",
			m.response.Manufacturer,
			m.response.Model,
			m.response.MAC,
			m.response.serialNumber,
			m.response.Software,
		)

		for i := range m.inputsConfiguration {
			b.WriteString(m.inputsConfiguration[i].View())
			if i < len(m.inputsConfiguration)-1 {
				b.WriteRune('\n')
			}
		}

		button := &blurredButton
		if m.focusIndexConfiguration == len(m.inputsConfiguration) {
			button = &focusedButton
		}
		fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	case response:
		if m.err != nil {
			fmt.Fprintf(&b, "erro: %+v • [Ctrl+C] ou ESC para reconfigurar acesso\n\n", m.err)
			return b.String()
		}
		fmt.Fprintf(&b, "[%+v] • [Ctrl+C] ou ESC para reconfigurar acesso\n\n", m.response)
	}

	return b.String()
}
