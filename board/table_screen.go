package board

import (
	"github.com/ascii-arcade/knuckle-bones/screen"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tableScreen struct {
	model *Model
	style lipgloss.Style

	selectedColumn int
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

	}

	return s.model, nil
}

func (s *tableScreen) View() string {
	mainPanelStyle := lipgloss.NewStyle().Width(s.model.width).Height(s.model.height).Align(lipgloss.Center, lipgloss.Center)

	boardStyle := lipgloss.NewStyle().
		Align(lipgloss.Center, lipgloss.Center).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#fff")).
		Width(33).
		Height(17)

	boardPlayerStyle := lipgloss.NewStyle().
		Align(lipgloss.Center, lipgloss.Center).
		Width(33).
		Height(17)

	boardTop := boardStyle.Render(
		s.model.Game.PlayerOneBoard.Render(s.selectedColumn),
	)

	boardBottom := boardStyle.Render(
		s.model.Game.PlayerTwoBoard.Render(-1),
	)

	pOneBoard := boardPlayerStyle.Height(33).Render("TEST\n0000")

	mainPanel := mainPanelStyle.Render(
		lipgloss.JoinHorizontal(
			lipgloss.Center,
			lipgloss.PlaceVertical(
				s.model.height,
				lipgloss.Bottom,
				lipgloss.JoinVertical(
					lipgloss.Bottom,
					pOneBoard,
				),
			),
			lipgloss.PlaceVertical(
				s.model.height,
				lipgloss.Center,
				lipgloss.JoinVertical(
					lipgloss.Bottom,
					boardBottom,
					boardTop,
				),
			),
			lipgloss.PlaceVertical(
				s.model.height,
				lipgloss.Top,
				lipgloss.JoinVertical(
					lipgloss.Bottom,
					pOneBoard,
				),
			),
		),
	)

	return mainPanel
}
