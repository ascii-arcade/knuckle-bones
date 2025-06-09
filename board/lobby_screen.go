package board

import (
	"fmt"

	"github.com/ascii-arcade/game-template/colors"
	"github.com/ascii-arcade/game-template/keys"
	"github.com/ascii-arcade/game-template/screen"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type lobbyScreen struct {
	model *Model
	style lipgloss.Style
}

func (m *Model) newLobbyScreen() *lobbyScreen {
	return &lobbyScreen{
		model: m,
		style: m.style,
	}
}

func (s *lobbyScreen) WithModel(model any) screen.Screen {
	s.model = model.(*Model)
	return s
}

func (s *lobbyScreen) Update(msg tea.Msg) (any, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		s.model.height, s.model.width = msg.Height, msg.Width
		return s.model, nil

	case tea.KeyMsg:
		if keys.LobbyStartGame.TriggeredBy(msg.String()) {
			if s.model.Player.IsHost() {
				_ = s.model.Game.Begin()
			}
		}
	}

	return s.model, nil
}

func (s *lobbyScreen) View() string {
	style := s.style.Width(s.model.width / 2)

	footer := s.model.lang().Get("board", "waiting_for_start")
	if s.model.Player.IsHost() {
		footer = fmt.Sprintf(s.model.lang().Get("board", "press_to_start"), keys.LobbyStartGame.String(s.style))

		if err := s.model.Game.IsPlayerCountOk(); err != nil {
			errorMessage := s.model.lang().Get("error", err.Error())
			footer = s.style.Foreground(colors.Error).Render(errorMessage)
		}
	}
	footer += "\n"
	footer += fmt.Sprintf(s.model.lang().Get("global", "quit"), keys.ExitApplication.String(s.style))

	header := s.model.Game.Code
	playerList := s.style.Render(s.playerList())

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		style.Align(lipgloss.Center).MarginBottom(2).Render(header),
		style.Render(playerList),
		style.Render(footer),
	)

	return s.style.Width(s.model.width).Height(s.model.height).Render(
		lipgloss.Place(
			s.model.width,
			s.model.height,
			lipgloss.Center,
			lipgloss.Center,
			s.style.
				Padding(2, 2).
				BorderStyle(lipgloss.NormalBorder()).
				Render(content),
		),
	)
}

func (s *lobbyScreen) playerList() string {
	playerList := ""
	for _, p := range s.model.Game.OrderedPlayers() {
		playerList += "* " + p.Name
		if p.Name == s.model.Player.Name {
			playerList += fmt.Sprintf(" (%s)", s.model.lang().Get("board", "player_list_you"))
		}
		if p.IsHost() {
			playerList += fmt.Sprintf(" (%s)", s.model.lang().Get("board", "player_list_host"))
		}
		playerList += "\n"
	}
	return playerList
}
