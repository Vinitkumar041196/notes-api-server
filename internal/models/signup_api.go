package models

type SignupRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string
}
