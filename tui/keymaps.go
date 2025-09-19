package tui

import (
	"github.com/charmbracelet/bubbles/key"
)

type listKeyMap struct {
	toggleSpinner    key.Binding
	toggleTitleBar   key.Binding
	toggleStatusBar  key.Binding
	togglePagination key.Binding
	toggleHelpMenu   key.Binding
	insertItem       key.Binding
}

func backroundListKeyMap() *listKeyMap {
	return &listKeyMap{
		insertItem: key.NewBinding(
			key.WithKeys("a"),
			key.WithHelp("a", "add item"),
		),
		toggleSpinner: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "toggle spinner"),
		),
		toggleTitleBar: key.NewBinding(
			key.WithKeys("T"),
			key.WithHelp("T", "toggle title"),
		),
		toggleStatusBar: key.NewBinding(
			key.WithKeys("S"),
			key.WithHelp("S", "toggle status"),
		),
		togglePagination: key.NewBinding(
			key.WithKeys("P"),
			key.WithHelp("P", "toggle pagination"),
		),
		toggleHelpMenu: key.NewBinding(
			key.WithKeys("H"),
			key.WithHelp("H", "toggle help"),
		),
	}
}

type modalKeyMap struct {
	submit   key.Binding
	cancel   key.Binding
	quit     key.Binding
	editItem key.Binding
}

func modalKeyMaps() *modalKeyMap {
	return &modalKeyMap{
		submit: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "submit"),
		),
		cancel: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "cancel"),
		),
		quit: key.NewBinding(
			key.WithKeys("ctrl+c"),
			key.WithHelp("ctrl+c", "close modal"),
		),
		// editItem: key.NewBinding(
		// 	key.WithKeys("e"),
		// 	key.WithHelp("e", "edit item"),
		// ),
	}
}

type delegateKeyMap struct {
	choose     key.Binding
	delete     key.Binding
	done       key.Binding
	inProgress key.Binding
	review     key.Binding
	notStarted key.Binding
}

func newDelegateKeyMap() *delegateKeyMap {
	return &delegateKeyMap{
		choose: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "choose"),
		),
		delete: key.NewBinding(
			key.WithKeys("ctrl+x", "backspace"),
			key.WithHelp("ctrl+x/backspace", "delete"),
		),
		done: key.NewBinding(
			key.WithKeys("d"),
			key.WithHelp("d", "mark as done"),
		),
		inProgress: key.NewBinding(
			key.WithKeys("p"),
			key.WithHelp("p", "mark in progress"),
		),
		review: key.NewBinding(
			key.WithKeys("r"),
			key.WithHelp("r", "mark as review"),
		),
		notStarted: key.NewBinding(
			key.WithKeys("n"),
			key.WithHelp("n", "mark not started"),
		),
	}
}
