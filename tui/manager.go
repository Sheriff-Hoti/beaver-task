package tui

import (
	"github.com/Sheriff-Hoti/beaver-task/database"
	overlay "github.com/Sheriff-Hoti/beaver-task/overlay"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type sessionState int

const (
	mainView sessionState = iota
	modalView
)

// Manager implements tea.Model, and manages the browser UI.
type Manager struct {
	state        sessionState
	windowWidth  int
	windowHeight int
	Foreground   tea.Model
	// Background   list.Model
	overlay      tea.Model
	list         list.Model
	queries      *database.Queries // Database queries for task management
	keys         *listKeyMap
	delegateKeys *delegateKeyMap
}

// Init initialises the Manager on program load. It partly implements the tea.Model interface.
func (m *Manager) Init() tea.Cmd {

	return nil
}

// Update handles event and manages internal state. It partly implements the tea.Model interface.
func (m *Manager) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.WindowSizeMsg:
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height
		h, v := appStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	// case tea.KeyMsg:
	// 	switch msg.String() {
	// 	case "q", "esc":
	// 		return m, tea.Quit

	// 	case " ":
	// 		if m.state == mainView {
	// 			m.state = modalView
	// 		} else {
	// 			m.state = mainView
	// 		}
	// 		return m, nil
	// 	}
	// }
	case tea.KeyMsg:
		// Don't match any of the keys below if we're actively filtering.
		if m.list.FilterState() == list.Filtering {
			break
		}

		switch {
		case key.Matches(msg, m.keys.toggleSpinner):
			cmd := m.list.ToggleSpinner()
			return m, cmd

		case key.Matches(msg, m.keys.toggleTitleBar):
			v := !m.list.ShowTitle()
			m.list.SetShowTitle(v)
			m.list.SetShowFilter(v)
			m.list.SetFilteringEnabled(v)
			return m, nil

		case key.Matches(msg, m.keys.toggleStatusBar):
			m.list.SetShowStatusBar(!m.list.ShowStatusBar())
			return m, nil

		case key.Matches(msg, m.keys.togglePagination):
			m.list.SetShowPagination(!m.list.ShowPagination())
			return m, nil

		case key.Matches(msg, m.keys.toggleHelpMenu):
			m.list.SetShowHelp(!m.list.ShowHelp())
			return m, nil

		case key.Matches(msg, m.keys.insertItem):
			m.delegateKeys.remove.SetEnabled(true)
			if m.state == mainView {
				m.state = modalView
			} else {
				m.state = mainView
			}
			// newItem := &Item{
			// 	TaskTitle:       "New Task",
			// 	TaskDescription: "This is a new task",
			// }
			// //add the insert logic here
			// insCmd := m.list.InsertItem(0, newItem)
			// statusCmd := m.list.NewStatusMessage(statusMessageStyle("Added " + newItem.Title()))
			// return m, tea.Batch(insCmd, statusCmd)
			return m, nil
		}
	}

	// This will also call our delegate's update function.
	newListModel, cmd := m.list.Update(message)
	m.list = newListModel

	fg, fgCmd := m.Foreground.Update(message)
	m.Foreground = fg

	// ls, bgCmd := m.list.Update(message)
	// m.list = ls

	cmds := []tea.Cmd{}
	cmds = append(cmds, cmd, fgCmd)

	return m, tea.Batch(cmds...)
}

// View applies and styling and handles rendering the view. It partly implements the tea.Model
// interface.
func (m *Manager) View() string {
	if m.state == modalView {
		return m.overlay.View()
	}
	return appStyle.Render(m.list.View())

}

func NewManager(queries *database.Queries, data_list []list.Item) *Manager {
	var (
		delegateKeys = newDelegateKeyMap()
		listKeys     = newListKeyMap()
	)

	fg := NewForeground()
	bg := NewBackground(data_list)

	return &Manager{
		state:      mainView,
		Foreground: fg,
		// Background: tasks,
		overlay: overlay.New(
			fg,
			bg,
			overlay.Center,
			overlay.Center,
			0,
			0,
		),
		queries:      queries,
		keys:         listKeys,
		delegateKeys: delegateKeys,
		list:         bg,
	}
}

func NewBackground(data_list []list.Item) list.Model {
	var (
		delegateKeys = newDelegateKeyMap()
		listKeys     = newListKeyMap()
	)

	delegate := newItemDelegate(delegateKeys)

	new_list := list.New(data_list, delegate, 0, 0)

	new_list.Title = "Tasks"
	new_list.Styles.Title = titleStyle

	new_list.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			listKeys.toggleSpinner,
			listKeys.insertItem,
			listKeys.toggleTitleBar,
			listKeys.toggleStatusBar,
			listKeys.togglePagination,
			listKeys.toggleHelpMenu,
		}
	}
	return new_list
}
