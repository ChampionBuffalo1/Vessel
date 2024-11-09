package charm

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Up    key.Binding
	Down  key.Binding
	Quit  key.Binding
	Enter key.Binding
}

var KeyMap = keyMap{
	Enter: key.NewBinding(
		key.WithKeys("enter"),
	),
	Up: key.NewBinding(
		key.WithKeys("k", "up"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp("↓/j", "move down"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "esc"),
		key.WithHelp("esc/ctrl+c", "quit"),
	),
}
