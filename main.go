package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle  = focusedStyle.Copy()
	noStyle      = lipgloss.NewStyle()

	focusedButton = focusedStyle.Copy().Render("[ OK ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("OK"))
)

// type CameraCfg struct {
// 	IPAddr       string
// 	Gateway      string
// 	SubnetMask   string
// 	ChannelTitle string
// 	Hostname     string
// 	Ponto        string
// }

type model struct {
	inputsAccess     []textinput.Model
	inputsCamera     []textinput.Model
	focusIndexAccess int
	focusIndexCamera int

	stage stage

	user string
	pass string
	addr string

	response CameraInfo

	err error // TODO definir Cmd proprio erro igual no github?
}

type stage int

const (
	access stage = iota
	camera
	response
)

func initialModel() model {
	// TODO ler base
	m := model{
		inputsAccess: make([]textinput.Model, 3),
		inputsCamera: make([]textinput.Model, 6),
	}

	var t textinput.Model
	for i := range m.inputsAccess {
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
			t.Placeholder = "usuário (admin)" // trocar para root quando for axis

		case 2:
			t.Placeholder = "senha"
			t.EchoMode = textinput.EchoPassword
			t.EchoCharacter = '•'
		}
		m.inputsAccess[i] = t
	}

	for i := range m.inputsCamera {
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
		m.inputsCamera[i] = t
	}

	return m
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.stage {
	case access:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "ctrl+c", "esc":
				return m, tea.Quit

			case "tab", "shift+tab", "enter", "up", "down":
				s := msg.String()

				// Pressed OK button save to model and goes to camera config
				if s == "enter" && m.focusIndexAccess == len(m.inputsAccess) {
					m.saveToModel()
					return m, nil
				}

				if s == "up" || s == "shift+tab" {
					m.focusIndexAccess--
				} else {
					m.focusIndexAccess++
				}

				if m.focusIndexAccess > len(m.inputsAccess) {
					m.focusIndexAccess = 0
				} else if m.focusIndexAccess < 0 {
					m.focusIndexAccess = len(m.inputsAccess)
				}

				cmds := make([]tea.Cmd, len(m.inputsAccess))
				for i := 0; i <= len(m.inputsAccess)-1; i++ {
					if i == m.focusIndexAccess {
						cmds[i] = m.inputsAccess[i].Focus()
						m.inputsAccess[i].PromptStyle = focusedStyle
						m.inputsAccess[i].TextStyle = focusedStyle
						continue
					}
					// remove focused state
					m.inputsAccess[i].Blur()
					m.inputsAccess[i].PromptStyle = noStyle
					m.inputsAccess[i].TextStyle = noStyle
				}

				return m, tea.Batch(cmds...)
			}
		}

		// Handle character input and blinking
		cmd := m.updateInputsAccess(msg)

		return m, cmd

	case camera:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "ctrl+c", "esc":
				m.cameraToAccess()
				return m, nil

			case "tab", "shift+tab", "enter", "up", "down":
				s := msg.String()

				// Pressed OK button save to model and goes to camera config
				if s == "enter" && m.focusIndexCamera == len(m.inputsCamera) {
					// m.saveToModel() // TODO dispatch urls
					// m.printerer()
					m.getCamera()
					return m, nil
				}

				if s == "up" || s == "shift+tab" {
					m.focusIndexCamera--
				} else {
					m.focusIndexCamera++
				}

				if m.focusIndexCamera > len(m.inputsCamera) {
					m.focusIndexCamera = 0
				} else if m.focusIndexCamera < 0 {
					m.focusIndexCamera = len(m.inputsCamera)
				}

				cmds := make([]tea.Cmd, len(m.inputsCamera))
				for i := 0; i <= len(m.inputsCamera)-1; i++ {
					if i == m.focusIndexCamera {
						cmds[i] = m.inputsCamera[i].Focus()
						m.inputsCamera[i].PromptStyle = focusedStyle
						m.inputsCamera[i].TextStyle = focusedStyle
						continue
					}
					// remove focused state
					m.inputsCamera[i].Blur()
					m.inputsCamera[i].PromptStyle = noStyle
					m.inputsCamera[i].TextStyle = noStyle
				}

				return m, tea.Batch(cmds...)
			}
		}

		// Handle character input and blinking
		cmd := m.updateInputsCamera(msg)

		return m, cmd
	}

	return m, nil
}

func (m *model) updateInputsAccess(msg tea.Msg) tea.Cmd {
	var cmds = make([]tea.Cmd, len(m.inputsAccess))

	// Only text inputsAccess with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputsAccess {
		m.inputsAccess[i], cmds[i] = m.inputsAccess[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m *model) updateInputsCamera(msg tea.Msg) tea.Cmd {
	var cmds = make([]tea.Cmd, len(m.inputsCamera))

	// Only text inputsCamera with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputsCamera {
		m.inputsCamera[i], cmds[i] = m.inputsCamera[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m *model) saveToModel() {
	m.addr = m.inputsAccess[0].Value()
	m.user = m.inputsAccess[1].Value()
	m.pass = m.inputsAccess[2].Value()

	m.stage = camera
	m.focusIndexAccess = 0
	m.inputsCamera[0].Focus()
	m.inputsCamera[0].PromptStyle = focusedStyle
	m.inputsCamera[0].TextStyle = focusedStyle
}

func (m *model) cameraToAccess() {
	m.stage = access
	m.focusIndexCamera = 0
	m.inputsAccess[0].Focus()
	m.inputsAccess[0].PromptStyle = focusedStyle
	m.inputsAccess[0].TextStyle = focusedStyle
}

func (m *model) getCamera() {
	camCfg, err := IdentificadorDeModelo(
		m.inputsAccess[0].Value(),
		m.inputsAccess[1].Value(),
		m.inputsAccess[2].Value(),
	)
	if err != nil {
		m.err = err
		return
	}

	m.response = camCfg
	m.stage = response
}

func (m model) View() string {
	var b strings.Builder

	b.WriteString("\nC O N F I G U R A D O R  -  D E  -  C Â M E R A\n")
	b.WriteString("Filipe Andrade -- filipe.engenhaira42@gmail.com\n\n")

	switch m.stage {
	case access:
		b.WriteString("[Ctrl+C] ou ESC para sair\n\n")

		for i := range m.inputsAccess {
			b.WriteString(m.inputsAccess[i].View())
			if i < len(m.inputsAccess)-1 {
				b.WriteRune('\n')
			}
		}

		button := &blurredButton
		if m.focusIndexAccess == len(m.inputsAccess) {
			button = &focusedButton
		}
		fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	case camera:
		fmt.Fprintf(&b, "[%s - %s] • [Ctrl+C] ou ESC para reconfigurar acesso\n\n", m.addr, m.user)

		for i := range m.inputsCamera {
			b.WriteString(m.inputsCamera[i].View())
			if i < len(m.inputsCamera)-1 {
				b.WriteRune('\n')
			}
		}

		button := &blurredButton
		if m.focusIndexCamera == len(m.inputsCamera) {
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

func main() {
	if err := tea.NewProgram(initialModel()).Start(); err != nil {
		fmt.Printf("erro: %s\n", err)
		os.Exit(1)
	}
}
