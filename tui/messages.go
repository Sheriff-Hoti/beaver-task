package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Custom message
type ItemChosenMsg struct {
	Value string
}

type ViewState struct {
	State viewState
}

type AddItemMsg struct {
	Value string
}

type StatusMessageTimeoutMsg struct{}

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

func addItemCmd(item string) tea.Cmd {
	return func() tea.Msg {
		return AddItemMsg{Value: item}
	}
}
