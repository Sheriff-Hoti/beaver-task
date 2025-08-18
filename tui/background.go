package tui

// import (
// 	"github.com/charmbracelet/bubbles/list"
// 	tea "github.com/charmbracelet/bubbletea"
// 	"github.com/charmbracelet/lipgloss"
// )

// var (
// 	appStyle = lipgloss.NewStyle().Padding(1, 2)

// 	titleStyle = lipgloss.NewStyle().
// 			Foreground(lipgloss.Color("#FFFDF5")).
// 			Background(lipgloss.Color("#25A065")).
// 			Padding(0, 1)

// 	statusMessageStyle = lipgloss.NewStyle().
// 				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
// 				Render
// )

// // Manager implements tea.Manager, and manages the browser UI.
// type Background struct {
// 	windowWidth  int
// 	windowHeight int
// 	list         *list.Model
// }

// // Init initialises the Manager on program load. It partly implements the tea.Manager interface.
// func (m *Background) Init() tea.Cmd {
// 	return nil
// }

// // Update handles event and manages internal state. It partly implements the tea.Manager interface.
// func (m *Background) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	// var cmds []tea.Cmd

// 	switch msg := msg.(type) {
// 	case tea.WindowSizeMsg:
// 		h, v := appStyle.GetFrameSize()
// 		m.list.SetSize(msg.Width-h, msg.Height-v)

// 		// case tea.KeyMsg:
// 		// 	// Don't match any of the keys below if we're actively filtering.
// 		// 	if m.list.FilterState() == list.Filtering {
// 		// 		break
// 		// 	}

// 		// 	switch {
// 		// 	case key.Matches(msg, m.keys.toggleSpinner):
// 		// 		cmd := m.list.ToggleSpinner()
// 		// 		return m, cmd

// 		// 	case key.Matches(msg, m.keys.toggleTitleBar):
// 		// 		v := !m.list.ShowTitle()
// 		// 		m.list.SetShowTitle(v)
// 		// 		m.list.SetShowFilter(v)
// 		// 		m.list.SetFilteringEnabled(v)
// 		// 		return m, nil

// 		// 	case key.Matches(msg, m.keys.toggleStatusBar):
// 		// 		m.list.SetShowStatusBar(!m.list.ShowStatusBar())
// 		// 		return m, nil

// 		// 	case key.Matches(msg, m.keys.togglePagination):
// 		// 		m.list.SetShowPagination(!m.list.ShowPagination())
// 		// 		return m, nil

// 		// 	case key.Matches(msg, m.keys.toggleHelpMenu):
// 		// 		m.list.SetShowHelp(!m.list.ShowHelp())
// 		// 		return m, nil

// 		// 	case key.Matches(msg, m.keys.insertItem):
// 		// 		m.delegateKeys.remove.SetEnabled(true)
// 		// 		newItem := &Item{
// 		// 			TaskTitle:       "New Task",
// 		// 			TaskDescription: "This is a new task",
// 		// 		}
// 		// 		//add the insert logic here
// 		// 		insCmd := m.list.InsertItem(0, newItem)
// 		// 		statusCmd := m.list.NewStatusMessage(statusMessageStyle("Added " + newItem.Title()))
// 		// 		return m, tea.Batch(insCmd, statusCmd)
// 		// 	}
// 	}

// 	// This will also call our delegate's update function.
// 	// newListModel, cmd := m.list.Update(msg)
// 	// m.list = newListModel
// 	// cmds = append(cmds, cmd)

// 	return m, nil
// }

// // View applies and styling and handles rendering the view. It partly implements the tea.Manager
// // interface.
// func (m *Background) View() string {

// 	return appStyle.Render(m.list.View())

// }
