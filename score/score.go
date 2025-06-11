package score

import "github.com/ascii-arcade/knucklebones/dice"

func Calculate(pool []dice.DicePool) int {
	score := 0

	for _, col := range pool {
		faceDupes := map[int]int{}
		for _, die := range col {
			faceDupes[die]++
		}

		for face, count := range faceDupes {
			switch count {
			case 1:
				score += face
			case 2:
				score += (face * count) * 2
			case 3:
				score += (face * count) * 3
			}
		}
	}

	return score
}
