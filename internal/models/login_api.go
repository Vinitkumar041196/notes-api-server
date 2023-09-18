package models

type LoginRequest struct {
	Email    string `json:"email"`
	Password string
}

type LoginSuccessResponse struct {
	SID string `json:"sid"`
}
