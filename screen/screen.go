package screen

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Screen interface {
	Update(tea.Msg) (any, tea.Cmd)
	View() string
	WithModel(any) Screen
}
