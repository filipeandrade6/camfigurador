package main

import (
	"fmt"
	"strings"
)

func (m model) View() string {
	var b strings.Builder

	b.WriteString("\nC O N F I G U R A D O R  -  D E  -  C Â M E R A\n")
	b.WriteString("Filipe Andrade -- filipe.engenhaira42@gmail.com\n\n")

	switch m.stage {
	case credentials:
		b.WriteString("[Ctrl+C] ou ESC para sair\n\n")

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
		// TODO colocar a resposta do acesso e remover o [Ctrl+C] ou ESC para reconfigurar acesso.
		if m.err != nil {
			fmt.Fprintf(&b, "erro de acesso: %s \n[Ctrl+C] ou ESC para reconfigurar acesso", m.err)
			return b.String()
		}

		fmt.Fprintf(&b, "Fabricante: %s\n         Modelo: %s\n            MAC: %s\nNúmero de série: %s\n       Software: %s\n\n",
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
