package messages

import (
	"github.com/ascii-arcade/knucklebones/games"
	"github.com/ascii-arcade/knucklebones/screen"
)

type (
	SwitchToMenuMsg  struct{}
	SwitchToBoardMsg struct{ Game *games.Game }
	SwitchScreenMsg  struct{ Screen screen.Screen }
	RefreshBoard     struct{}
	RollMsg          struct{}
)
