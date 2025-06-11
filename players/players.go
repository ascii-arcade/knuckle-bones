package players

import (
	"context"

	"github.com/ascii-arcade/knucklebones/dice"
	"github.com/ascii-arcade/knucklebones/generaterandom"
	"github.com/ascii-arcade/knucklebones/language"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
)

var players = make(map[string]*Player)

func NewPlayer(ctx context.Context, sess ssh.Session, langPref *language.LanguagePreference) *Player {
	var player *Player

	defer func() {
		go func() {
			<-player.ctx.Done()
			player.Connected = false
			for _, fn := range player.onDisconnect {
				fn()
			}
		}()
	}()

	var exists bool
	player, exists = players[sess.User()]
	if exists {
		player.UpdateChan = make(chan struct{})
		player.Connected = true
		player.isHost = false
		player.ctx = ctx

		return player
	}

	board := make([]dice.DicePool, 3)
	for i := range board {
		board[i] = make(dice.DicePool, 3)
	}

	player = &Player{
		Name:               generaterandom.Name(langPref.Lang),
		Score:              0,
		UpdateChan:         make(chan struct{}),
		Board:              board,
		Pool:               make(dice.DicePool, 1),
		Color:              lipgloss.Color(generaterandom.Color()),
		LanguagePreference: langPref,
		Sess:               sess,
		Connected:          true,
		ctx:                ctx,
	}
	players[sess.User()] = player

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
