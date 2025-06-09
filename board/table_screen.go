package board

import (
	"fmt"
	"strings"

	"github.com/ascii-arcade/game-template/keys"
	"github.com/ascii-arcade/game-template/screen"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tableScreen struct {
	model *Model
	style lipgloss.Style
}

func (m *Model) newTableScreen() *tableScreen {
	return &tableScreen{
		model: m,
		style: m.style,
	}
}

func (s *tableScreen) WithModel(model any) screen.Screen {
	s.model = model.(*Model)
	return s
}

func (s *tableScreen) Update(msg tea.Msg) (any, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		s.model.height, s.model.width = msg.Height, msg.Width
		return s.model, nil

	case tea.KeyMsg:
		if keys.GameIncrementPoint.TriggeredBy(msg.String()) {
			s.model.Game.Count(s.model.Player)
		}
	}

	return s.model, nil
}

func (s *tableScreen) View() string {
	disconnectedPlayers := s.model.Game.GetDisconnectedPlayers()
	if len(disconnectedPlayers) > 0 {
		var names []string
		for _, p := range disconnectedPlayers {
			names = append(names, p.Name)
		}
		return s.style.Render(
			lipgloss.JoinVertical(
				lipgloss.Center,
				s.model.style.Align(lipgloss.Center).MarginBottom(2).Render(s.model.Game.Code),
				fmt.Sprintf(s.model.lang().Get("board", "disconnected_player"), strings.Join(names, ", ")),
				s.style.Render(fmt.Sprintf(s.model.lang().Get("global", "quit"), keys.ExitApplication.String(s.style))),
			),
		)
	}

	counts := ""
	for _, p := range s.model.Game.OrderedPlayers() {
		counts += fmt.Sprintf("%s: %d\n", p.Name, p.Count)
	}

	return s.style.Render(fmt.Sprintf(s.model.lang().Get("board", "you_are"), s.model.Player.Name)) +
		"\n\n" + counts +
		"\n\n" + s.style.Render(fmt.Sprintf(s.model.lang().Get("global", "quit"), keys.ExitApplication.String(s.style)))
}
