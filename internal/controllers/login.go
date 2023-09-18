package controllers

import (
	"net/http"
	"notes-api-server/internal/models"
	"notes-api-server/internal/utils"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	UserService models.UserService
}

func (s *LoginController) LoginHandler(c *gin.Context) {
	req := models.LoginRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "invalid request format", Code: http.StatusBadRequest})
		return
	}

	user, err := s.UserService.GetUser(req.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error(), Code: http.StatusBadRequest})
		return
	}

	if utils.GetHashedString(req.Password) != user.Password {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "username and password doesn't match", Code: http.StatusUnauthorized})
		return
	}

	sid := s.UserService.CreateUserSession(req.Email)

	c.JSON(http.StatusOK, models.LoginSuccessResponse{SID: sid})
}
