package board

import (
	"time"

	"github.com/ascii-arcade/knuckle-bones/keys"
	"github.com/ascii-arcade/knuckle-bones/messages"
	"github.com/ascii-arcade/knuckle-bones/screen"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tableScreen struct {
	model *Model
	style lipgloss.Style

	rollTickCount int
	rolling       bool
}

const (
	rollFrames   = 15
	rollInterval = 200 * time.Millisecond
)

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

	case messages.RollMsg:
		if !s.model.game.IsTurn(s.model.player) {
			return s.model, nil
		}

		if s.rollTickCount < rollFrames {
			s.rollTickCount++
			s.model.game.RollDice(s.rolling)
			return s.model, tea.Tick(rollInterval, func(time.Time) tea.Msg {
				return messages.RollMsg{}
			})
		}
		s.rolling = false
		s.model.game.RollDice(s.rolling)

	case tea.KeyMsg:
		if !s.model.game.IsTurn(s.model.player) {
			return s.model, nil
		}

		if keys.ActionRoll.TriggeredBy(msg.String()) {
			if !s.model.game.Rolled() && !s.rolling {
				s.rollTickCount = 0
				s.rolling = true
				return s.model, tea.Tick(rollInterval, func(time.Time) tea.Msg {
					return messages.RollMsg{}
				})
			}
		}
	}

	return s.model, nil
}

func (s *tableScreen) View() string {
	mainPanelStyle := lipgloss.NewStyle().Width(s.model.width).Height(s.model.height).Align(lipgloss.Center, lipgloss.Center)

	boardStyle := lipgloss.NewStyle().
		Align(lipgloss.Center, lipgloss.Center).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#fff")).
		Margin(0).
		Width(33).
		Height(17)

	boardPlayerStyle := lipgloss.NewStyle().
		Align(lipgloss.Center, lipgloss.Center).
		Margin(0).
		Width(33).
		Height(17)

	me := s.model.player
	them := s.model.game.GetOpponent(s.model.player)

	boardTop := boardStyle.Render(
		them.Board.Render(),
	)

	boardBottom := boardStyle.Height(16).AlignVertical(lipgloss.Bottom).Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			me.Board.Render(),
			lipgloss.JoinHorizontal(
				lipgloss.Center,
				lipgloss.PlaceHorizontal(9, lipgloss.Center, "1"),
				lipgloss.PlaceHorizontal(9, lipgloss.Center, "2"),
				lipgloss.PlaceHorizontal(9, lipgloss.Center, "3"),
			),
		),
	)

	theirBoard := boardPlayerStyle.Height(33).Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			them.Name,
			"0",
			them.Pool.Render(),
		),
	)

	myBoard := boardPlayerStyle.Height(33).Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			me.Pool.Render(),
			me.Name,
			"0",
		),
	)

	mainPanel := mainPanelStyle.Render(
		lipgloss.JoinHorizontal(
			lipgloss.Center,
			lipgloss.PlaceVertical(
				s.model.height,
				lipgloss.Bottom,
				lipgloss.JoinVertical(
					lipgloss.Bottom,
					myBoard,
				),
			),
			lipgloss.PlaceVertical(
				s.model.height,
				lipgloss.Center,
				lipgloss.JoinVertical(
					lipgloss.Bottom,
					boardTop,
					boardBottom,
				),
			),
			lipgloss.PlaceVertical(
				s.model.height,
				lipgloss.Top,
				lipgloss.JoinVertical(
					lipgloss.Bottom,
					theirBoard,
				),
			),
		),
	)

	return mainPanel
}
