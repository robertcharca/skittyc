package tui

import (
    "fmt"
)

func (tuim terminalUIModel) View() string {
    s := "What would you like to do?\n"

    for i, choice := range tuim.choices {
        cursor := " "
        if tuim.cursor == i {
            cursor = ">"
        }

        checked := " "
        if _, ok := tuim.selected[i]; ok {
            checked = "x"
        }

        s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
    }

    s += "\nPress q to quit"

    return s
}
