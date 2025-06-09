package keys

import (
	"slices"

	"github.com/charmbracelet/lipgloss"
)

type Keys []string

func (k Keys) TriggeredBy(msg string) bool {
	return slices.Contains(k, msg)
}

func (k Keys) String(style lipgloss.Style) string {
	return k.IndexedString(0, style)
}

func (k Keys) IndexedString(index int, style lipgloss.Style) string {
	if len(k) == 0 {
		return ""
	}
	return style.Bold(true).Italic(true).Render("'" + k[index] + "'")
}

var (
	MenuJoinGame     = Keys{"j"}
	MenuStartNewGame = Keys{"n"}
	MenuEnglish      = Keys{"1"}
	MenuSpanish      = Keys{"2"}

	PreviousScreen = Keys{"esc"}
	Submit         = Keys{"enter"}

	ExitApplication    = Keys{"ctrl+c"}
	GameIncrementPoint = Keys{"a"}
	LobbyStartGame     = Keys{"s"}
)
