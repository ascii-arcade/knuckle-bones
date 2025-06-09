package menu

import (
	"github.com/ascii-arcade/game-template/messages"
	"github.com/ascii-arcade/game-template/screen"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const splashLogo = `++------------------------------------------------------------------------------++
++------------------------------------------------------------------------------++
||                                                                              ||
||                                                                              ||
||      _    ____   ____ ___ ___        _    ____   ____    _    ____  _____    ||
||     / \  / ___| / ___|_ _|_ _|      / \  |  _ \ / ___|  / \  |  _ \| ____|   ||
||    / _ \ \___ \| |    | | | |_____ / _ \ | |_) | |     / _ \ | | | |  _|     ||
||   / ___ \ ___) | |___ | | | |_____/ ___ \|  _ <| |___ / ___ \| |_| | |___    ||
||  /_/   \_\____/ \____|___|___|   /_/   \_\_| \_\\____/_/   \_\____/|_____|   ||
||                                                                              ||
||                                                                              ||
||                                                                              ||
++------------------------------------------------------------------------------++
++------------------------------------------------------------------------------++`

type splashScreen struct {
	model *Model
	style lipgloss.Style
}

func (m *Model) newSplashScreen() *splashScreen {
	return &splashScreen{
		model: m,
		style: m.style,
	}
}

func (s *splashScreen) WithModel(model any) screen.Screen {
	s.model = model.(*Model)
	return s
}

func (s *splashScreen) Update(msg tea.Msg) (any, tea.Cmd) {
	switch msg.(type) {
	case doneMsg:
		return s.model, func() tea.Msg {
			return messages.SwitchScreenMsg{
				Screen: s.model.newOptionScreen(),
			}
		}
	}
	return s.model, nil
}

func (s *splashScreen) View() string {
	style := s.style.
		Width(s.model.width).
		Height(s.model.height)

	return style.Render(
		lipgloss.Place(
			s.model.width,
			s.model.height,
			lipgloss.Center,
			lipgloss.Center,
			splashLogo,
		),
	)
}
