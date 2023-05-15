package tui

import "github.com/charmbracelet/bubbles/list"

type mainModel struct {
	list     list.Model
	choice   string
	quitting bool
}
