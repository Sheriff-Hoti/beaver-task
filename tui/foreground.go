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

func (m Foreground) appErrorBoundaryView(text string) string {
	return lipgloss.PlaceHorizontal(
		m.windowWidth,
		lipgloss.Left,
		lipgloss.NewStyle().
			Foreground(indigo).
			Bold(true).
			Padding(0, 1, 0, 2).Render(text),
		lipgloss.WithWhitespaceChars("/"),
		lipgloss.WithWhitespaceForeground(red),
	)
}

func (m Foreground) errorView() string {
	var s string
	for _, err := range m.form.Errors() {
		s += err.Error()
	}
	return s
}

// Update handles event and manages internal state. It partly implements the tea.Model interface.
func (m *Foreground) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := message.(type) {
	case tea.WindowSizeMsg:
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height

	case ViewState:
		m.state = msg.State
		if m.state == modalView {
			m.form = NewForm()
			return m, m.form.Init()
		} else {
			m.form = nil
		}
		// if m.state == modalView {
		// 	// focus the form when modal opens
		// 	m.form = NewForm()
		// 	cmds = append(cmds, m.form.Init())
		// }

	case tea.KeyMsg:

		if m.state == mainView {
			// If we're in modal view, don't process any key events.
			return m, nil
		}
		switch {

		case key.Matches(msg, m.keys.cancel):
			return m, tea.Batch(changeViewState(mainView))
		// case key.Matches(msg, m.keys.submit):
		// 	m.form = NewForm()
		// 	// return m, tea.Batch(changeViewState(mainView), addItemCmd("niiice"))
		// 	return m, nil

		case key.Matches(msg, m.keys.editItem):
			return m, nil
		case key.Matches(msg, m.keys.quit):
			return m, tea.Quit
		}
	}

	if m.form == nil {
		return m, tea.Batch(cmds...)
	}

	//if focus issue run form init eveytime the modal opens
	form, cmd := m.form.Update(message)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	if m.form.State == huh.StateCompleted {
		// Quit when the form is done.
		title := m.form.GetString("title")
		cmds = append(cmds, changeViewState(mainView), addItemCmd(title))
		m.form = nil
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
		Width(m.windowWidth / 2).
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
	if m.form == nil {
		return ""
	}

	return foreStyle.Render(lipgloss.JoinVertical(lipgloss.Left, title.Render("Create a Task"), m.form.View()))
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
	title_input := huh.NewInput().Title("Title").Prompt("> ").Key("title").Validate(huh.ValidateNotEmpty())
	form := huh.NewForm(
		huh.NewGroup(
			title_input,
		).WithShowHelp(true),
	)
	return form
}
