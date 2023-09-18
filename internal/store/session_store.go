package store

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type SessionStore struct {
	store map[string]string
	mux   sync.Mutex
}

func NewSessionStore() *SessionStore {
	s := new(SessionStore)
	s.store = make(map[string]string)
	return s
}

func (s *SessionStore) AddSession(email string) string {
	sid := uuid.New().String()
	s.mux.Lock()
	s.store[sid] = email
	s.mux.Unlock()
	return sid
}

func (s *SessionStore) GetSession(sid string) (string, error) {
	email, found := s.store[sid]
	if !found {
		return "", fmt.Errorf("invalid session")
	}
	return email, nil
}
