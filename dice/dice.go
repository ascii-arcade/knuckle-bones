package dice

import (
	"math/rand/v2"
	"slices"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type DicePool []int

func (p *DicePool) Roll() {
	for i := range *p {
		(*p)[i] = rand.IntN(6) + 1
	}
}

func (p *DicePool) Contains(face int) bool {
	return slices.Contains(*p, face)
}

func (p *DicePool) Add(face int) {
	*p = append(*p, face)
}

func (p *DicePool) Remove(face int) bool {
	for i, n := range *p {
		if n == face {
			*p = slices.Delete(*p, i, i+1)
			return true
		}
	}
	return false
}

func (p *DicePool) RenderCharacters() string {
	if len(*p) == 0 {
		return ""
	}

	output := ""
	for i, n := range *p {
		output += diceCharacters[n]
		if i != len(*p)-1 {
			output += " "
		}
	}

	return strings.TrimSpace(output)
}

func (p *DicePool) Render() string {
	diceCount := len(*p)

	gridSize := 3
	grid := make([][]string, gridSize)
	for col := range grid {
		grid[col] = make([]string, gridSize)
		for row := range grid[col] {
			idx := col*gridSize + row
			grid[col][row] = ""

			if idx < diceCount {
				grid[col][row] = lipgloss.NewStyle().Render(dieFaces[(*p)[idx]])
			}
		}
	}

	columns := make([]string, gridSize)
	for i := range grid {
		columns[i] = lipgloss.JoinVertical(lipgloss.Top, grid[i]...)
	}
	return lipgloss.JoinHorizontal(lipgloss.Center, columns...)
}
