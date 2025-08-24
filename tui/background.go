package tui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// Manager implements tea.Manager, and manages the browser UI.
type Background struct {
	windowWidth  int
	windowHeight int
	list         list.Model
	keys         *listKeyMap
	delegateKeys *delegateKeyMap
	state        viewState
}

// Init initialises the Manager on program load. It partly implements the tea.Manager interface.
func (m *Background) Init() tea.Cmd {
	return nil
}

// Update handles event and manages internal state. It partly implements the tea.Manager interface.
func (m *Background) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := message.(type) {
	case tea.WindowSizeMsg:
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height
		h, v := appStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case ViewState:
		m.state = msg.State

	case AddItemMsg:
		m.list.InsertItem(0, &Task{TaskTitle: msg.Value, TaskDescription: "This is a new task", TaskCreatedAt: "a few seconds ago"})
		m.list.NewStatusMessage(statusMessageStyle("Added " + msg.Value))

	case tea.KeyMsg:

		if m.state == modalView {
			// If we're in modal view, don't process any key events.
			return m, nil
		}

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

			// m.delegateKeys.remove.SetEnabled(true)
			// newItem := &Task{
			// 	TaskTitle:       "New Task",
			// 	TaskDescription: "This is a new taskaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			// }
			// //add the insert logic here
			// insCmd := m.list.InsertItem(0, newItem)
			// statusCmd := m.list.NewStatusMessage(statusMessageStyle("Added " + newItem.Title()))
			// return m, tea.Batch(insCmd, statusCmd)
			return m, tea.Batch(changeViewState(modalView))
		}
	}

	newListModel, cmd := m.list.Update(message)
	m.list = newListModel
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)

}

// View applies and styling and handles rendering the view. It partly implements the tea.Manager
// interface.
func (m *Background) View() string {
	return appStyle.Render(m.list.View())

}

func NewBackground(list_items []list.Item) *Background {
	var (
		delegateKeys = newDelegateKeyMap()
		listKeys     = backroundListKeyMap()
	)

	delegate := NewCustomDelegate(delegateKeys)

	tasks := list.New(list_items, delegate, 0, 0)
	tasks.SetShowStatusBar(false)
	tasks.SetFilteringEnabled(true)
	tasks.SetShowHelp(true)

	tasks.Title = "Tasks"
	tasks.Styles.Title = titleStyle

	tasks.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			listKeys.toggleSpinner,
			listKeys.insertItem,
			listKeys.toggleTitleBar,
			listKeys.toggleStatusBar,
			listKeys.togglePagination,
			listKeys.toggleHelpMenu,
		}
	}

	return &Background{
		list:         tasks,
		keys:         listKeys,
		delegateKeys: delegateKeys,
	}
}
