package service

import (
	"notes-api-server/internal/app"
	"notes-api-server/internal/models"
	"notes-api-server/internal/store"
	"notes-api-server/internal/utils"
)

type UserService struct {
	UserStore    *store.UserStore
	SessionStore *store.SessionStore
}

func NewUserService(app *app.App) models.UserService {
	return &UserService{
		UserStore: app.UserStore,
		SessionStore: app.SessionStore,
	}
}

func (us *UserService) AddUser(u models.User) error {
	u.Password = utils.GetHashedString(u.Password)
	return us.UserStore.AddUser(u)
}

func (us *UserService) GetUser(email string) (models.User, error) {
	return us.UserStore.GetUser(email)
}

func (us *UserService) CreateUserSession(email string) string {
	return us.SessionStore.AddSession(email)
}
