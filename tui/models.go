package tui

import "github.com/charmbracelet/bubbles/list"

type fontModel struct {
	list     list.Model
	choice   string
	quitting bool
}
