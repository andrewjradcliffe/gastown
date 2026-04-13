package convoy

import "github.com/charmbracelet/bubbles/key"

// KeyMap defines the key bindings for the convoy TUI.
type KeyMap struct {
	Up       key.Binding
	Down     key.Binding
	PageUp   key.Binding
	PageDown key.Binding
	Top      key.Binding
	Bottom   key.Binding
	Toggle   key.Binding // expand/collapse
	Help     key.Binding
	Quit     key.Binding
}

// DefaultKeyMap returns the default key bindings.
func DefaultKeyMap() KeyMap {
	return KeyMap{
		Up: key.NewBinding(
			key.WithKeys("up", "k", "ctrl+p"),
			key.WithHelp("↑/k/ctrl+p", "up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down", "j", "ctrl+n"),
			key.WithHelp("↓/j/ctrl+n", "down"),
		),
		PageUp: key.NewBinding(
			key.WithKeys("pgup", "ctrl+u", "alt+v"),
			key.WithHelp("pgup/alt+v", "page up"),
		),
		PageDown: key.NewBinding(
			key.WithKeys("pgdown", "ctrl+d", "ctrl+v"),
			key.WithHelp("pgdn/ctrl+v", "page down"),
		),
		Top: key.NewBinding(
			key.WithKeys("home", "g", "alt+<"),
			key.WithHelp("g/alt+<", "top"),
		),
		Bottom: key.NewBinding(
			key.WithKeys("end", "G", "alt+>"),
			key.WithHelp("G/alt+>", "bottom"),
		),
		Toggle: key.NewBinding(
			key.WithKeys("enter", " "),
			key.WithHelp("enter/space", "expand/collapse"),
		),
		Help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "help"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", "esc", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
	}
}

// ShortHelp returns keybindings to show in the help view.
func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.Toggle, k.Quit, k.Help}
}

// FullHelp returns keybindings for the expanded help view.
func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.PageUp, k.PageDown},
		{k.Top, k.Bottom, k.Toggle},
		{k.Help, k.Quit},
	}
}
