package tui

import (
	"context"

	"github.com/Sheriff-Hoti/beaver-task/database"
	"github.com/Sheriff-Hoti/beaver-task/overlay"
	tea "github.com/charmbracelet/bubbletea"
)

type viewState int

const (
	mainView viewState = iota
	modalView
)

// Manager implements tea.Model, and manages the browser UI.
type Manager struct {
	state        viewState
	windowWidth  int
	windowHeight int
	foreground   tea.Model
	background   tea.Model
	overlay      tea.Model
	queries      *database.Queries
	ctx          context.Context
}

// Init initialises the Manager on program load. It partly implements the tea.Model interface.
func (m *Manager) Init() tea.Cmd {
	// m.state = mainView
	// m.foreground = &Foreground{}
	// m.background = &Background{}
	// m.overlay = overlay.New(
	// 	m.foreground,
	// 	m.background,
	// 	overlay.Center,
	// 	overlay.Center,
	// 	0,
	// 	0,
	// )
	return tea.Batch(
		m.foreground.Init(),
		m.background.Init(),
	)
}

// Update handles event and manages internal state. It partly implements the tea.Model interface.
func (m *Manager) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	cmds := []tea.Cmd{}

	switch msg := message.(type) {
	case tea.WindowSizeMsg:
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height

	case AddItemMsg:
		m.queries.CreateTask(m.ctx, database.CreateTaskParams{
			Title: msg.Value,
		})

	case ItemChosenMsg:
		// m.state = modalView

	case DeleteItemMsg:
		m.queries.DeleteTaskByTitle(m.ctx, msg.Title)

	case ViewState:
		m.state = msg.State
		// if m.state == modalView {
		// 	// reset & focus the form when opening the modal
		// 	fg, fgCmd := m.foreground.Update(ResetFormMsg{})
		// 	m.foreground = fg
		// 	cmds = append(cmds, fgCmd)
		// }
	}

	fg, fgCmd := m.foreground.Update(message)
	m.foreground = fg

	bg, bgCmd := m.background.Update(message)
	m.background = bg

	cmds = append(cmds, fgCmd, bgCmd)

	return m, tea.Batch(cmds...)
}

// View applies and styling and handles rendering the view. It partly implements the tea.Model
// interface.
func (m *Manager) View() string {
	if m.state == modalView {
		return m.overlay.View()
	}
	return m.background.View()
}

func NewManager(tasks []database.Task, queries *database.Queries, ctx context.Context) *Manager {

	foreground := NewForeground()
	bacground := NewBackground(fromDatabaseTasks(tasks))
	overlay := overlay.New(
		foreground,
		bacground,
		overlay.Center,
		overlay.Center,
		0,
		0,
	)

	return &Manager{
		state:      mainView,
		foreground: foreground,
		background: bacground,
		overlay:    overlay,
		queries:    queries,
		ctx:        ctx,
	}
}
