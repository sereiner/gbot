package session

import (
	"github.com/sereiner/gbot/gbot"
	"github.com/sereiner/library/utility"
)

type Session struct {
	gbot gbot.Gbot
}

func NewSession(gbot gbot.Gbot) *Session {
	return &Session{gbot: gbot}
}

func (s *Session) randomKey() string {
	return utility.GetGUID()[:8]
}

func (s *Session) CurrentSession() string {
	argSession := s.gbot.Session()
	if len(argSession) == 0 {
		argSession = s.randomKey()
	}
	return argSession
}

func (s *Session) Has(session string) bool {
	return false
}
