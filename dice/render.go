package dice

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	diceCharacters = map[int]string{
		1: "⚀",
		2: "⚁",
		3: "⚂",
		4: "⚃",
		5: "⚄",
		6: "⚅",
	}

	dieStyle = lipgloss.NewStyle().
			Width(7).
			Height(3).
			Align(lipgloss.Center).
			Border(lipgloss.Border(lipgloss.RoundedBorder())).
			MarginRight(1)

	dieFaces = map[int]string{
		1: dieStyle.Render("\n" + center("●")),
		2: dieStyle.Render(strings.Join([]string{left("●"), "", right("●")}, "\n")),
		3: dieStyle.Render(strings.Join([]string{left("●"), center("●"), right("●")}, "\n")),
		4: dieStyle.Render(strings.Join([]string{center("●   ●"), "", center("●   ●")}, "\n")),
		5: dieStyle.Render(strings.Join([]string{center("●   ●"), center("●"), center("●   ●")}, "\n")),
		6: dieStyle.Render(strings.Join([]string{center("●   ●"), center("●   ●"), center("●   ●")}, "\n")),
	}
)

func left(s string) string {
	return lipgloss.NewStyle().Width(5).Align(lipgloss.Left).Render(s)
}

func center(s string) string {
	return lipgloss.NewStyle().Width(5).Align(lipgloss.Center).Render(s)
}

func right(s string) string {
	return lipgloss.NewStyle().Width(5).Align(lipgloss.Right).Render(s)
}

func GetDieCharacter(face int) string {
	if face < 1 || face > 6 {
		return " "
	}
	return diceCharacters[face]
}
