package tui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type delegateKeyMap struct {
	choose key.Binding
	delete key.Binding
}

// newItemDelegate returns a list.DefaultDelegate used to render and handle events for an individual
// list item (Task). It wires per-item key handling (choose/remove) and provides short/full help bindings.
func newItemDelegate(keys *delegateKeyMap) CustomDelegate {
	d := NewCustomDelegate()

	d.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
		var title string

		if i, ok := m.SelectedItem().(*Task); ok {
			title = i.Title()
		} else {
			return nil
		}

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, keys.choose):
				return func() tea.Msg {
					return ItemChosenMsg{Value: title}
				}
			case key.Matches(msg, keys.delete):
				index := m.Index()
				m.RemoveItem(index)
				if len(m.Items()) == 0 {
					keys.delete.SetEnabled(false)
				}
				return tea.Batch(m.NewStatusMessage(statusMessageStyle("Deleted "+title)), deleteItemCmd(title))
			}
		}

		return nil
	}

	help := []key.Binding{keys.choose, keys.delete}

	d.ShortHelpFunc = func() []key.Binding {
		return help
	}

	d.FullHelpFunc = func() [][]key.Binding {
		return [][]key.Binding{help}
	}

	return d
}

// Additional short help entries. This satisfies the help.KeyMap interface and
// is entirely optional.
func (d delegateKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		d.choose,
		d.delete,
	}
}

// Additional full help entries. This satisfies the help.KeyMap interface and
// is entirely optional.
func (d delegateKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			d.choose,
			d.delete,
		},
	}
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
	}
}
