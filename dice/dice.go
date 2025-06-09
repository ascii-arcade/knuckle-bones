package dice

import (
	"math/rand/v2"
	"slices"
	"strings"

	"github.com/ascii-arcade/knuckle-bones/score"
	"github.com/charmbracelet/lipgloss"
)

type DicePool []int

func NewDicePool(size int) DicePool {
	p := make(DicePool, size)
	for i := range p {
		p[i] = 1
	}
	return p
}

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

func (p *DicePool) Score() (int, error) {
	return score.Calculate([][]int{}, false)
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

func (p *DicePool) Render(start int, end int) string {
	diceCount := len(*p)
	if diceCount == 0 {
		return ""
	}
	if end > diceCount {
		end = diceCount
	}
	if start >= end {
		return ""
	}

	topCount := (diceCount + 1) / 2
	bottomCount := diceCount / 2

	topDice := make([]string, 0)
	for i := range topCount {
		topDice = append(topDice, dieFaces[(*p)[i]])
	}

	bottomDice := make([]string, 0)
	for i := range bottomCount {
		bottomDice = append(bottomDice, dieFaces[(*p)[i+topCount]])
	}

	return lipgloss.JoinVertical(
		lipgloss.Center,
		lipgloss.JoinHorizontal(lipgloss.Top, topDice...),
		lipgloss.JoinHorizontal(lipgloss.Top, bottomDice...),
	)
}
