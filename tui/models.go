package tui

import (
    //tea "github.com/charmbracelet/bubbletea"
)

type terminalUIModel struct { 
    choices []string
    cursor int
    selected map[int]struct{}
}
