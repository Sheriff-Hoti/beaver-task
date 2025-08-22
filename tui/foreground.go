package tui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	huh "github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

// Model implements tea.Model, and manages the browser UI.
type Foreground struct {
	windowWidth  int
	windowHeight int
	state        viewState
	form         *huh.Form
	keys         *modalKeyMap
}

// Init initialises the Model on program load. It partly implements the tea.Model interface.
func (m *Foreground) Init() tea.Cmd {
	return m.form.Init()
}

// Update handles event and manages internal state. It partly implements the tea.Model interface.
func (m *Foreground) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	// if m.form.State == huh.StateCompleted {
	// 	m.form = NewForm()
	// 	return m, tea.Batch(changeViewState(mainView))
	// }

	switch msg := message.(type) {
	case tea.WindowSizeMsg:
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height

	case ViewState:
		m.state = msg.State

	case tea.KeyMsg:

		if m.state == mainView {
			// If we're in modal view, don't process any key events.
			return m, nil
		}
		switch {

		case key.Matches(msg, m.keys.cancel):
			return m, tea.Batch(changeViewState(mainView))
		case key.Matches(msg, m.keys.submit):
			m.form = NewForm()
			return m, tea.Batch(changeViewState(mainView), addItemCmd("niiice"))
		case key.Matches(msg, m.keys.editItem):
			return m, nil
		case key.Matches(msg, m.keys.quit):
			return m, tea.Quit
		}
	}
	//if focus issue run form init eveytime the modal opens
	form, cmd := m.form.Update(message)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

// View applies and styling and handles rendering the view. It partly implements the tea.Model
// interface.
func (m *Foreground) View() string {
	foreStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder(), true).
		BorderForeground(lipgloss.Color("6")).
		Padding(0, 1).
		Width(m.windowWidth / 3).
		Height(m.windowHeight / 3)
		// .MarginLeft(m.windowWidth / 4).MarginRight(m.windowWidth / 4)

	// boldStyle := lipgloss.NewStyle().Bold(true)
	// title := boldStyle.Render("Bubble Tea Overlay")
	// content := "Hello! I'm in a modal window.\n\nPress <space> to close the window."
	// layout := lipgloss.JoinVertical(lipgloss.Left, title, content)
	// if m.form.State == huh.StateCompleted {
	// 	title := m.form.GetString("title")
	// 	return fmt.Sprintf("You selected: %s,", title)
	// }

	return foreStyle.Render(m.form.View())
}

func NewForeground() *Foreground {

	return &Foreground{
		windowWidth:  0,
		windowHeight: 0,
		state:        mainView,
		form:         NewForm(),
		keys:         modalKeyMaps(),
	}
}

func NewForm() *huh.Form {
	title_input := huh.NewInput().Title("Title").Prompt(">").Key("title")
	return huh.NewForm(
		huh.NewGroup(
			title_input,
		),
	)
}
