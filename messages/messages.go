package messages

import (
	"github.com/ascii-arcade/game-template/games"
	"github.com/ascii-arcade/game-template/screen"
)

type (
	SwitchToMenuMsg  struct{}
	SwitchToBoardMsg struct{ Game *games.Game }
	SwitchScreenMsg  struct{ Screen screen.Screen }
	RefreshBoard     struct{}
)
