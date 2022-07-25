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

type CameraCfg struct {
    IPAddr       string
    Gateway      string
    SubnetMask   string
    ChannelTitle string
    Hostname     string
    Ponto        string
}

type model struct {
    inputsAccess     []textinput.Model
    inputsCamera     []textinput.Model
    focusIndexAccess int
    focusIndexCamera int

    stage stage

    user         string
    pass         string
    addr         string
    manufacturer string
}

type stage int

const (
    access stage = iota
    camera
)

func initialModel() model {
    // TODO ler base
    m := model{
        inputsAccess: make([]textinput.Model, 4),
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
            t.Placeholder = "fabricante (dahua|axis)"

        case 2:
            t.Placeholder = "usuário (admin)" // trocar para root quando for axis

        case 3:
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
        // TODO considerar outros casos (mouse, etc)
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
            m.printerer()
            return m, nil // TODO voltar para access?

        case "tab", "shift+tab", "enter", "up", "down":
            s := msg.String()

            // Pressed OK button save to model and goes to camera config
            if s == "enter" && m.focusIndexCamera == len