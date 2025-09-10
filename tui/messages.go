package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Custom message
type ItemChosenMsg struct {
	Value string
}

type DeleteItemMsg struct {
	Title string
}

type ViewState struct {
	State viewState
}

type AddItemResultMsg struct {
	task *Task
	err  error
}

type AddItemRequestMsg struct {
	title string
}

func chooseItemCmd(val string) tea.Cmd {
	return func() tea.Msg {
		return ItemChosenMsg{Value: val}
	}
}

func deleteItemCmd(title string) tea.Cmd {
	return func() tea.Msg {
		return DeleteItemMsg{Title: title}
	}
}

func changeViewState(state viewState) tea.Cmd {
	return func() tea.Msg {
		return ViewState{State: state}
	}
}

func addItemResultCmd(task *Task, err error) tea.Cmd {
	return func() tea.Msg {
		return AddItemResultMsg{task: task, err: err}
	}
}

func addItemRequestCmd(title string) tea.Cmd {
	return func() tea.Msg {
		return AddItemRequestMsg{title: title}
	}
}
