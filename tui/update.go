package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (tuim terminalUIModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String(){
        case "ctrl+c", "q":
            return tuim, tea.Quit
                
        case "up", "k":
            if tuim.cursor > 0 {
                tuim.cursor-- 
            }

        case "down", "j":
            if tuim.cursor < len(tuim.choices)-1 {
                tuim.cursor++
            }

        case "enter", " ":
            _, ok := tuim.selected[tuim.cursor]
            if ok {
                delete(tuim.selected, tuim.cursor)    
            } else {
                tuim.selected[tuim.cursor] = struct{}{}        
            }

        }
    }

    return tuim, nil
}
