package tui

import tea "github.com/charmbracelet/bubbletea"

// Custom message
type ItemChosenMsg struct {
	Value string
}

type ViewState struct {
	State viewState
}

func chooseItemCmd(val string) tea.Cmd {
	return func() tea.Msg {
		return ItemChosenMsg{Value: val}
	}
}

func changeViewState(state viewState) tea.Cmd {
	return func() tea.Msg {
		return ViewState{State: state}
	}
}
