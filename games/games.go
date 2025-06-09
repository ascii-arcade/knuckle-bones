package games

import (
	"errors"

	"github.com/ascii-arcade/game-template/generaterandom"
)

var (
	ErrGameInProgress = errors.New("game_already_in_progress")
	ErrGameNotFound   = errors.New("game_not_found")
)

var games = make(map[string]*Game)

func New() *Game {
	game := &Game{
		Code:    generaterandom.Code(),
		players: make([]*Player, 0),
	}
	games[game.Code] = game

	return game
}

func GetOpenGame(code string) (*Game, error) {
	game, exists := games[code]
	if !exists {
		return nil, ErrGameNotFound
	}
	if game.inProgress {
		return game, ErrGameInProgress
	}

	return game, nil
}

func GetAll() map[string]*Game {
	return games
}
