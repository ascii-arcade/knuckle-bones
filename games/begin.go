package games

import "errors"

func (s *Game) Begin() error {
	return s.withErrLock(func() error {
		if error := s.IsPlayerCountOk(); error != nil {
			return error
		}

		s.inProgress = true
		return nil
	})
}

func (s *Game) IsPlayerCountOk() error {
	if s.PlayerTwo == nil {
		return errors.New("not_enough_players")
	}
	return nil
}
