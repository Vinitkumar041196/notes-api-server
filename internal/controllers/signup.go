package controllers

import (
	"net/http"
	"notes-api-server/internal/models"

	"github.com/gin-gonic/gin"
)

type SignupController struct {
	UserService models.UserService
}

func (s *SignupController) SignupHandler(c *gin.Context) {
	req := models.SignupRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "invalid request format", Code: http.StatusBadRequest})
		return
	}

	user := models.User(req)

	err = s.UserService.AddUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error(), Code: http.StatusBadRequest})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{Message: "success"})
}
