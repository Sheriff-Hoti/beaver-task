package tui

//got this code from: https://github.com/charmbracelet/bubbles/blob/master/list/defaultitem.go

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
)

// CustomItemStyles defines styling for a default list item.
// See CustomItemView for when these come into play.
type CustomItemStyles struct {
	// The Normal state.
	NormalTitle lipgloss.Style
	NormalDesc  lipgloss.Style

	// The selected item state.
	SelectedTitle lipgloss.Style
	SelectedDesc  lipgloss.Style

	// The dimmed state, for when the filter input is initially activated.
	DimmedTitle lipgloss.Style
	DimmedDesc  lipgloss.Style

	// Characters matching the current filter, if any.
	FilterMatch lipgloss.Style

	// CreatedAt timestamp styles.
	NormalCreatedAt   lipgloss.Style
	SelectedCreatedAt lipgloss.Style
	DimmedCreatedAt   lipgloss.Style
}

// NewCustomItemStyles returns style definitions for a default item. See
// CustomItemView for when these come into play.
func NewCustomItemStyles() (s CustomItemStyles) {
	s.NormalTitle = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#dddddd"}).
		Padding(0, 0, 0, 2) //nolint:mnd

	s.NormalDesc = s.NormalTitle.
		Foreground(lipgloss.AdaptiveColor{Light: "#A49FA5", Dark: "#777777"})

	s.SelectedTitle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#AD58B4"}).
		Foreground(lipgloss.AdaptiveColor{Light: "#EE6FF8", Dark: "#EE6FF8"}).
		Padding(0, 0, 0, 1)

	s.SelectedDesc = s.SelectedTitle.
		Foreground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#AD58B4"})

	s.DimmedTitle = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#A49FA5", Dark: "#777777"}).
		Padding(0, 0, 0, 2) //nolint:mnd

	s.DimmedDesc = s.DimmedTitle.
		Foreground(lipgloss.AdaptiveColor{Light: "#C2B8C2", Dark: "#4D4D4D"})

	s.FilterMatch = lipgloss.NewStyle().Underline(true)

	// CreatedAt styles (default to desc-like appearance)
	s.NormalCreatedAt = lipgloss.NewStyle().
		Background(lipgloss.Color("#44475a")). // soft dark bg
		Foreground(lipgloss.Color("#f8f8f2")). // light text
		Padding(0, 1).                         // some breathing room
		Bold(true).
		Align(lipgloss.Right)
	s.SelectedCreatedAt = lipgloss.NewStyle().
		Background(lipgloss.Color("#bd93f9")). // purple bg
		Foreground(lipgloss.Color("#1e1e2e")). // dark text
		Padding(0, 1).
		Bold(true).
		Align(lipgloss.Right)
	s.DimmedCreatedAt = lipgloss.NewStyle().
		Background(lipgloss.Color("#282a36")). // very dark gray
		Foreground(lipgloss.Color("#6272a4")). // muted text
		Padding(0, 1).
		Align(lipgloss.Right)

	return s
}

// CustomItem describes an item designed to work with CustomDelegate.
type CustomItem interface {
	*Task
	Title() string
	Description() string
	FilterValue() string
	CreatedAt() string
}

// CustomDelegate is a standard delegate designed to work in lists. It's
// styled by CustomItemStyles, which can be customized as you like.
//
// The description line can be hidden by setting Description to false, which
// renders the list as single-line-items. The spacing between items can be set
// with the SetSpacing method.
//
// Setting UpdateFunc is optional. If it's set it will be called when the
// ItemDelegate called, which is called when the list's Update function is
// invoked.
//
// Settings ShortHelpFunc and FullHelpFunc is optional. They can be set to
// include items in the list's default short and full help menus.
type CustomDelegate struct {
	ShowDescription bool
	Styles          CustomItemStyles
	UpdateFunc      func(tea.Msg, *list.Model) tea.Cmd
	ShortHelpFunc   func() []key.Binding
	FullHelpFunc    func() [][]key.Binding
	height          int
	spacing         int
}

// NewCustomDelegate creates a new delegate with default styles.
func NewCustomDelegate() CustomDelegate {
	const defaultHeight = 2
	const defaultSpacing = 1
	return CustomDelegate{
		ShowDescription: true,
		Styles:          NewCustomItemStyles(),
		height:          defaultHeight,
		spacing:         defaultSpacing,
	}
}

// SetHeight sets delegate's preferred height.
func (d *CustomDelegate) SetHeight(i int) {
	d.height = i
}

// Height returns the delegate's preferred height.
// This has effect only if ShowDescription is true,
// otherwise height is always 1.
func (d CustomDelegate) Height() int {
	if d.ShowDescription {
		return d.height
	}
	return 1
}

// SetSpacing sets the delegate's spacing.
func (d *CustomDelegate) SetSpacing(i int) {
	d.spacing = i
}

// Spacing returns the delegate's spacing.
func (d CustomDelegate) Spacing() int {
	return d.spacing
}

// Update checks whether the delegate's UpdateFunc is set and calls it.
func (d CustomDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	if d.UpdateFunc == nil {
		return nil
	}
	return d.UpdateFunc(msg, m)
}

// Render prints an item.
func (d CustomDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	var (
		title, desc, created_at string
		matchedRunes            []int
		s                       = &d.Styles
	)

	if i, ok := item.(*Task); ok {
		title = i.Title()
		desc = i.Description()
		created_at = i.CreatedAt()
	} else {
		return
	}

	if m.Width() <= 0 {
		// short-circuit
		return
	}

	// Prevent text from exceeding list width
	textwidth := m.Width() - s.NormalTitle.GetPaddingLeft() - s.NormalTitle.GetPaddingRight()
	title = ansi.Truncate(title, textwidth, ellipsis)

	// cap created_at so it can't overflow the line
	// maxCreated := 20
	// if textwidth < 20 {
	// 	maxCreated = textwidth / 2
	// 	if maxCreated < 5 {
	// 		maxCreated = 5
	// 	}
	// }
	// created_at = ansi.Truncate(created_at, maxCreated, ellipsis)

	if d.ShowDescription {
		var lines []string
		for i, line := range strings.Split(desc, "\n") {
			if i >= d.height-1 {
				break
			}
			lines = append(lines, ansi.Truncate(line, textwidth, ellipsis))
		}
		desc = strings.Join(lines, "\n")
	}

	// Conditions
	var (
		isSelected  = index == m.Index()
		emptyFilter = m.FilterState() == list.Filtering && m.FilterValue() == ""
		isFiltered  = m.FilterState() == list.Filtering || m.FilterState() == list.FilterApplied
	)

	if isFiltered && index < len(m.Items()) {
		// Get indices of matched characters
		matchedRunes = m.MatchesForItem(index)
	}

	if emptyFilter {
		title = s.DimmedTitle.Render(title)
		desc = s.DimmedDesc.Render(desc)
		created_at = s.DimmedCreatedAt.Render(created_at)

	} else if isSelected && m.FilterState() != list.Filtering {
		if isFiltered {
			// Highlight matches
			unmatched := s.SelectedTitle.Inline(true)
			matched := unmatched.Inherit(s.FilterMatch)
			title = lipgloss.StyleRunes(title, matchedRunes, matched, unmatched)
		}
		title = s.SelectedTitle.Render(title)
		desc = s.SelectedDesc.Render(desc)
		created_at = s.SelectedCreatedAt.Render(created_at)
	} else {
		if isFiltered {
			// Highlight matches
			unmatched := s.NormalTitle.Inline(true)
			matched := unmatched.Inherit(s.FilterMatch)
			title = lipgloss.StyleRunes(title, matchedRunes, matched, unmatched)
		}
		title = s.NormalTitle.Render(title)
		created_at = s.NormalCreatedAt.Render(created_at)
		desc = s.NormalDesc.Render(desc)
	}

	gap := max(m.Width()-lipgloss.Width(title)-lipgloss.Width(created_at), 1)

	if d.ShowDescription {
		line := lipgloss.JoinHorizontal(
			lipgloss.Top,
			title,
			strings.Repeat(" ", gap),
			created_at,
		)
		fmt.Fprintf(w, "%s\n%s", line, desc) //nolint: errcheck
		return
	}
	fmt.Fprintf(w, "%s", title) //nolint: errcheck
}

// ShortHelp returns the delegate's short help.
func (d CustomDelegate) ShortHelp() []key.Binding {
	if d.ShortHelpFunc != nil {
		return d.ShortHelpFunc()
	}
	return nil
}

// FullHelp returns the delegate's full help.
func (d CustomDelegate) FullHelp() [][]key.Binding {
	if d.FullHelpFunc != nil {
		return d.FullHelpFunc()
	}
	return nil
}
