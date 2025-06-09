package messages

import (
	"github.com/ascii-arcade/knuckle-bones/games"
	"github.com/ascii-arcade/knuckle-bones/screen"
)

type (
	SwitchToMenuMsg  struct{}
	SwitchToBoardMsg struct{ Game *games.Game }
	SwitchScreenMsg  struct{ Screen screen.Screen }
	RefreshBoard     struct{}
)
