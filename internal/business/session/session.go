package session

import (
	"Stas-sH/authWithSessions/pkg/utils"
)

type sessionData struct {
	Username string
}

type Session struct {
	data map[string]*sessionData
}

var InMemorySession *Session

func NewSession() *Session {
	s := new(Session)

	s.data = make(map[string]*sessionData)

	return s
}

func (s *Session) Init(userName string) string {
	sessionId := utils.GenerateId()

	data := &sessionData{Username: userName}
	s.data[sessionId] = data

	return sessionId
}

func (s *Session) GetInfo(sessionId string) string {
	data := s.data[sessionId]

	if data == nil {
		return ""
	}
	return data.Username
}
