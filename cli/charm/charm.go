package charm

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	PullPage page = iota
	ContainerImage
)

func NewBubbleTea(page page) {
	model := newModel(page)
	prog := tea.NewProgram(model)
	if _, err := prog.Run(); err != nil {
		panic(err)
	}
}

type model struct {
	width    int
	height   int
	page     uint8
	keymap   keyMap
	state    *state
	renderer *lipgloss.Renderer
}

func (m model) Init() tea.Cmd {
	switch m.page {
	case PullPage:
		return m.PullInit()
	}
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.Quit):
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	switch m.page {
	case PullPage:
		m, cmd = m.PullUpdate(msg)
	}
	return m, cmd
}

func (m model) View() string {
	switch m.page {
	case PullPage:
		return m.PullView()
	}
	return ""
}

type state struct {
	pull pullState
}

type page = uint8

func newModel(page page) model {
	m := model{
		page:     page,
		keymap:   keys,
		state:    newState(page),
		renderer: lipgloss.DefaultRenderer(),
	}
	return m
}

func newState(page page) *state {
	state := &state{}
	switch page {
	case PullPage:
		state.pull = newPullState()
	}
	return state
}
