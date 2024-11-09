package charm

import (
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type pullState struct {
	index      uint8
	searchTerm string
	list       list.Model
	help       help.Model
	input      textinput.Model
}

func newPullState() pullState {
	input := textinput.New()
	input.Focus()
	input.Width = 30

	return pullState{
		input: input,
		list:  list.New(nil, list.NewDefaultDelegate(), 0, 0),
		help:  help.New(),
	}
}

func (m model) PullInit() tea.Cmd {
	return textinput.Blink
}

func (m model) PullUpdate(msg tea.Msg) (model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.Enter):
			m.state.pull.searchTerm = m.state.pull.input.Value()
			m.state.pull.index = 0
		}
	}

	var cmd tea.Cmd
	if m.state.pull.searchTerm == "" {
		m.state.pull.input, cmd = m.state.pull.input.Update(msg)
	} else {
		m.state.pull.list, cmd = m.state.pull.list.Update(msg)
	}
	return m, cmd

}

func (m model) PullView() string {
	if m.state.pull.searchTerm == "" {
		return m.state.pull.input.View()
	}

	suggestionList := []list.Item{}
	m.state.pull.list.SetItems(suggestionList)
	m.state.pull.list.SetHeight(m.height - 1)
	m.state.pull.list.SetWidth(m.width)

	return fmt.Sprintf("%s\n\n%s", m.state.pull.list.View(), m.state.pull.help.View(m.keymap))
}
