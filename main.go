package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	inputsCredentials     []textinput.Model
	focusIndexCredentials int

	inputsConfiguration     []textinput.Model
	focusIndexConfiguration int

	stage stage

	response ConfigurationInfo

	err error // TODO definir Cmd proprio erro igual no github?
}

type stage int

const (
	credentials stage = iota
	configuration
	response
)

func initialModel() model {
	// TODO ler base
	m := model{
		inputsCredentials:   make([]textinput.Model, 3),
		inputsConfiguration: make([]textinput.Model, 6),
	}

	var t textinput.Model
	for i := range m.inputsCredentials {
		t = textinput.New()
		t.CursorStyle = cursorStyle
		t.CharLimit = 20

		switch i {
		case 0:
			t.Placeholder = "IP (192.168.1.108)"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle

		case 1:
			t.Placeholder = "usuário (admin)"

		case 2:
			t.Placeholder = "senha"
			t.EchoMode = textinput.EchoPassword
			t.EchoCharacter = '•'
		}
		m.inputsCredentials[i] = t
	}

	for i := range m.inputsConfiguration {
		t = textinput.New()
		t.CursorStyle = cursorStyle
		t.CharLimit = 50

		switch i {
		case 0:
			t.Placeholder = "IP"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle

		case 1:
			t.Placeholder = "Gateway"

		case 2:
			t.Placeholder = "Máscara de subrede"

		case 3:
			t.Placeholder = "Descrição no vídeo"

		case 4:
			t.Placeholder = "Ponto"

		case 5:
			t.Placeholder = "Patrimônio"
		}
		m.inputsConfiguration[i] = t
	}

	return m
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func main() {
	if err := tea.NewProgram(initialModel()).Start(); err != nil {
		fmt.Printf("erro: %s\n", err)
		os.Exit(1)
	}
}
