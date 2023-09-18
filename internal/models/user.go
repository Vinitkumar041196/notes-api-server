package models

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserService interface {
	AddUser(User) error
	GetUser(string) (User, error)
	CreateUserSession(string) string
}
