package games

import (
	"context"

	"github.com/ascii-arcade/game-template/language"
	"github.com/charmbracelet/ssh"
)

type Player struct {
	Name      string
	Count     int
	TurnOrder int

	isHost    bool
	connected bool

	UpdateChan         chan struct{}
	LanguagePreference *language.LanguagePreference

	Sess ssh.Session

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

func (p *Player) MakeHost() *Player {
	p.isHost = true
	return p
}

func (p *Player) IsHost() bool {
	return p.isHost
}

func (p *Player) OnDisconnect(fn func()) {
	p.onDisconnect = append(p.onDisconnect, fn)
}

func (p *Player) incrementCount() {
	p.Count++
}
