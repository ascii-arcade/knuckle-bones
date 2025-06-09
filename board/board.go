package board

import (
	"github.com/ascii-arcade/game-template/config"
	"github.com/ascii-arcade/game-template/games"
	"github.com/ascii-arcade/game-template/keys"
	"github.com/ascii-arcade/game-template/language"
	"github.com/ascii-arcade/game-template/messages"
	"github.com/ascii-arcade/game-template/screen"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	width  int
	height int
	screen screen.Screen
	style  lipgloss.Style

	Player *games.Player
	Game   *games.Game
}

func NewModel(width, height int, style lipgloss.Style, player *games.Player) Model {
	m := Model{
		width:  width,
		height: height,
		style:  style,
		Player: player,
	}

	m.screen = m.newTableScreen()
	return m
}

func (m Model) Init() tea.Cmd {
	return waitForRefreshSignal(m.Player.UpdateChan)
}

func (m *Model) lang() *language.Language {
	return m.Player.LanguagePreference.Lang
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.RefreshBoard:
		return m, waitForRefreshSignal(m.Player.UpdateChan)

	case tea.KeyMsg:
		if keys.ExitApplication.TriggeredBy(msg.String()) {
			m.Game.RemovePlayer(m.Player)
			return m, tea.Quit
		}
	}

	screenModel, cmd := m.activeScreen().Update(msg)
	return screenModel.(*Model), cmd
}

func (m Model) View() string {
	if m.width < config.MinimumWidth {
		return m.lang().Get("error", "window_too_narrow")
	}
	if m.height < config.MinimumHeight {
		return m.lang().Get("error", "window_too_short")
	}

	return m.activeScreen().View()
}

func (m *Model) activeScreen() screen.Screen {
	if m.Game.InProgress() {
		return m.newTableScreen()
	} else {
		return m.newLobbyScreen()
	}
}

func waitForRefreshSignal(ch chan struct{}) tea.Cmd {
	return func() tea.Msg {
		return messages.RefreshBoard(<-ch)
	}
}
