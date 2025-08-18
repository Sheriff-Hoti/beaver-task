package tui

import (
	"github.com/Sheriff-Hoti/beaver-task/database"
	"github.com/Sheriff-Hoti/beaver-task/overlay"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type sessionState int

const (
	mainView sessionState = iota
	modalView
)

type Task struct {
	ID              int64
	TaskTitle       string
	TaskDescription string
}

func (t Task) FilterValue() string { return t.TaskTitle }
func (t Task) Title() string       { return t.TaskTitle }
func (t Task) Description() string { return t.TaskDescription }

func fromDatabaseTask(task *database.Task) *Task {
	return &Task{
		ID:              task.ID,
		TaskTitle:       task.Title,
		TaskDescription: task.Description.String,
	}
}

func fromDatabaseTasks(tasks []database.Task) []list.Item {
	items := make([]list.Item, len(tasks))
	for i, task := range tasks {
		items[i] = fromDatabaseTask(&task)
	}
	return items
}

// Manager implements tea.Model, and manages the browser UI.
type Manager struct {
	state        sessionState
	windowWidth  int
	windowHeight int
	foreground   tea.Model
	background   tea.Model
	overlay      tea.Model
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
	return nil
}

// Update handles event and manages internal state. It partly implements the tea.Model interface.
func (m *Manager) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.WindowSizeMsg:
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height

	case ItemChosenMsg:
		m.state = modalView

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc":
			return m, tea.Quit

		case " ":
			if m.state == mainView {
				m.state = modalView
			} else {
				m.state = mainView
			}
			return m, nil

		case "a":
			if m.state == modalView {
				//this blocks the background from adding a new task
				return m, nil
			}

		}

	}

	fg, fgCmd := m.foreground.Update(message)
	m.foreground = fg

	bg, bgCmd := m.background.Update(message)
	m.background = bg

	cmds := []tea.Cmd{}
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

func NewManager(tasks []database.Task) *Manager {

	foreground := &Foreground{}
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
	}
}
