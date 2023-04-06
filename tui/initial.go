package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func initialModel() terminalUIModel {
    return terminalUIModel{
        choices: []string{"Change your font", "Change your color scheme"},
        selected: make(map[int]struct{}),
    }
}

func (tuim terminalUIModel) Init() tea.Cmd {
    return nil
}
