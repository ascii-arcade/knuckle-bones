package games

import (
	"sync"

	"github.com/ascii-arcade/knuckle-bones/dice"
	"github.com/ascii-arcade/knuckle-bones/players"
	"github.com/charmbracelet/ssh"
)

type Game struct {
	Code      string
	PlayerOne *players.Player
	PlayerTwo *players.Player

	turn   int
	rolled bool

	inProgress bool
	mu         sync.Mutex
}

func (g *Game) InProgress() bool {
	return g.inProgress
}

func (g *Game) refresh() {
	players := []*players.Player{g.PlayerOne, g.PlayerTwo}
	for _, p := range players {
		if p != nil && p.UpdateChan != nil {
			select {
			case p.UpdateChan <- struct{}{}:
			default:
			}
		}
	}
}

func (g *Game) Rolled() bool {
	r := false
	g.withLock(func() {
		r = g.rolled
	})
	return r
}

func (g *Game) withLock(fn func()) {
	g.mu.Lock()
	defer func() {
		g.refresh()
		g.mu.Unlock()
	}()
	fn()
}

func (g *Game) withErrLock(fn func() error) error {
	g.mu.Lock()
	defer func() {
		g.refresh()
		g.mu.Unlock()
	}()
	return fn()
}

func (g *Game) AddPlayer(player *players.Player) error {
	return g.withErrLock(func() error {
		if _, ok := g.getPlayer(player.Sess); ok {
			return nil
		}

		if g.inProgress {
			return ErrGameInProgress
		}

		player.OnDisconnect(func() {
			if !g.inProgress {
				g.RemovePlayer(player)
			}
		})

		if player.IsHost() {
			g.PlayerOne = player
		} else {
			g.PlayerTwo = player
		}

		return nil
	})
}

func (g *Game) RemovePlayer(player *players.Player) {
	g.withLock(func() {
		if player, exists := g.getPlayer(player.Sess); exists {
			close(player.UpdateChan)
			if g.PlayerOne == player {
				g.PlayerOne = nil
			} else if g.PlayerTwo == player {
				g.PlayerTwo = nil
			}
		}
	})
}

func (g *Game) getPlayer(sess ssh.Session) (*players.Player, bool) {
	if g.PlayerOne != nil && g.PlayerOne.Sess.User() == sess.User() {
		return g.PlayerOne, true
	} else if g.PlayerTwo != nil && g.PlayerTwo.Sess.User() == sess.User() {
		return g.PlayerTwo, true
	}
	return nil, false
}

func (g *Game) GetPlayers() []*players.Player {
	var players []*players.Player
	if g.PlayerOne != nil {
		players = append(players, g.PlayerOne)
	}
	if g.PlayerTwo != nil {
		players = append(players, g.PlayerTwo)
	}
	return players
}

func (g *Game) GetDisconnectedPlayers() []*players.Player {
	var players []*players.Player
	g.withLock(func() {
		if !g.PlayerOne.Connected {
			players = append(players, g.PlayerOne)
		}
		if !g.PlayerTwo.Connected {
			players = append(players, g.PlayerTwo)
		}
	})

	if len(players) == 2 {
		g.RemovePlayer(players[0])
		g.RemovePlayer(players[1])
		return nil
	}

	return players
}

func (g *Game) HasPlayer(player *players.Player) bool {
	_, exists := g.getPlayer(player.Sess)
	return exists
}

func (g *Game) nextTurn() {
	if g.turn == 0 {
		g.turn = 1
	} else {
		g.turn = 0
	}

	g.rolled = false
}

func (g *Game) RollDice(rolling bool) {
	g.withLock(func() {
		switch g.turn {
		case 0:
			g.PlayerOne.Pool.Roll()
		case 1:
			g.PlayerTwo.Pool.Roll()
		}

		if !rolling {
			g.rolled = true
		}
	})
}

func (g *Game) GetTurnPlayer() *players.Player {
	if g.turn == 0 {
		return g.PlayerOne
	}
	return g.PlayerTwo
}

func (g *Game) IsTurn(p *players.Player) bool {
	return g.GetTurnPlayer().Name == p.Name
}

func (g *Game) IsPlayerOne(p *players.Player) bool {
	return g.PlayerOne.Name == p.Name
}

func (g *Game) GetOpponent(p *players.Player) *players.Player {
	if g.PlayerOne == p {
		return g.PlayerTwo
	}
	return g.PlayerOne
}

func (g *Game) PlaceDie(p *players.Player, column int) error {
	return g.withErrLock(func() error {
		if !g.IsTurn(p) {
			return ErrNotYourTurn
		}

		if !g.rolled {
			return ErrDiceNotRolled
		}

		spot := nextSpot(p.Board[column])
		if spot == -1 {
			return ErrColumnFull
		}

		p.Board[column][spot] = p.Pool[0]
		p.Pool = make(dice.DicePool, 1)

		removeSame(g.GetOpponent(p).Board[column], p.Board[column][spot])

		if full(p.Board) {

		}

		g.nextTurn()
		return nil
	})
}

func nextSpot(pool dice.DicePool) int {
	for i, face := range pool {
		if face == 0 {
			return i
		}
	}
	return -1
}

func full(board []dice.DicePool) bool {
	for _, col := range board {
		if nextSpot(col) != -1 {
			return false
		}
	}
	return true
}

func removeSame(column dice.DicePool, face int) {
	for i, n := range column {
		if n == face {
			column[i] = 0 // Remove the die by setting it to 0
		}
	}
}
