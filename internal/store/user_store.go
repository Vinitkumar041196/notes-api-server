package store

import (
	"fmt"
	"notes-api-server/internal/models"
	"sync"
)

type UserStore struct {
	store map[string]models.User
	mux   sync.Mutex
}

func NewUserStore() *UserStore {
	u := new(UserStore)
	u.store = make(map[string]models.User)
	return u
}

func (u *UserStore) AddUser(user models.User) error {
	_, found := u.store[user.Email]
	if found {
		return fmt.Errorf("user already exists")
	}

	u.mux.Lock()
	u.store[user.Email] = user
	u.mux.Unlock()

	return nil
}

func (u *UserStore) GetUser(email string) (models.User, error) {
	user, found := u.store[email]
	if !found {
		return user, fmt.Errorf("user doesn't exists")
	}

	return user, nil
}
