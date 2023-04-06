package tui

import (
	"fmt"
	"os"
	tea "github.com/charmbracelet/bubbletea"
)

func Execute() {
    p := tea.NewProgram(initialModel())
    
    if _, err := p.Run(); err != nil {
        fmt.Printf("There's been an error: %v", err)
        os.Exit(1)
    }
}
