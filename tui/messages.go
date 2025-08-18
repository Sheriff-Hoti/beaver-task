package tui

import tea "github.com/charmbracelet/bubbletea"

// Custom message
type ItemChosenMsg struct {
	Value string
}

func chooseItemCmd(val string) tea.Cmd {
	return func() tea.Msg {
		return ItemChosenMsg{Value: val}
	}
}
