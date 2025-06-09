package games

func (s *Game) Count(player *Player) {
	s.withLock(func() {
		player.incrementCount()
	})
}
