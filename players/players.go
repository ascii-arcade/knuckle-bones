package players

import (
	"context"

	"github.com/ascii-arcade/knuckle-bones/dice"
	"github.com/ascii-arcade/knuckle-bones/generaterandom"
	"github.com/ascii-arcade/knuckle-bones/language"
	"github.com/charmbracelet/ssh"
)

var players = make(map[string]*Player)

func NewPlayer(ctx context.Context, sess ssh.Session, langPref *language.LanguagePreference) *Player {
	player, exists := players[sess.User()]
	if exists {
		player.UpdateChan = make(chan struct{})
		player.Connected = true
		player.isHost = false
		player.ctx = ctx

		goto RETURN
	}

	player = &Player{
		Name:               generaterandom.Name(langPref.Lang),
		Score:              0,
		UpdateChan:         make(chan struct{}),
		Board:              make(dice.DicePool, 9),
		Pool:               make(dice.DicePool, 1),
		LanguagePreference: langPref,
		Sess:               sess,
		Connected:          true,
		ctx:                ctx,
	}
	players[sess.User()] = player

RETURN:
	go func() {
		<-player.ctx.Done()
		player.Connected = false
		for _, fn := range player.onDisconnect {
			fn()
		}
	}()

	return player
}

func RemovePlayer(player *Player) {
	if _, exists := players[player.Sess.User()]; exists {
		close(player.UpdateChan)
		delete(players, player.Sess.User())
	}
}

func GetPlayerCount() int {
	return len(players)
}

func GetConnectedPlayerCount() int {
	count := 0
	for _, player := range players {
		if player.Connected {
			count++
		}
	}
	return count
}
