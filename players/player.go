package players

import (
	"context"

	"github.com/ascii-arcade/knuckle-bones/dice"
	"github.com/ascii-arcade/knuckle-bones/language"
	"github.com/charmbracelet/ssh"
)

type Player struct {
	Name      string
	Score     int
	TurnOrder int

	Connected bool

	UpdateChan         chan struct{}
	LanguagePreference *language.LanguagePreference

	Sess ssh.Session

	isHost bool

	Board []dice.DicePool
	Pool  dice.DicePool

	onDisconnect []func()
	ctx          context.Context
}

func (p *Player) SetName(name string) *Player {
	p.Name = name
	return p
}

func (p *Player) SetTurnOrder(order int) *Player {
	p.TurnOrder = order
	return p
}

func (p *Player) OnDisconnect(fn func()) {
	p.onDisconnect = append(p.onDisconnect, fn)
}

func (p *Player) MakeHost() {
	p.isHost = true
}

func (p *Player) IsHost() bool {
	return p.isHost
}
