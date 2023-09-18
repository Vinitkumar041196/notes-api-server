package controllers

import (
	"net/http"
	"notes-api-server/internal/models"

	"github.com/gin-gonic/gin"
)

type NotesController struct {
	NotesService models.NotesService
}

func (s *NotesController) GetAllNotesHandler(c *gin.Context) {
	req := models.GetAllNotesRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "invalid request format", Code: http.StatusBadRequest})
		return
	}

	notes, err := s.NotesService.GetAllNotes(req.SID.SID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: err.Error(), Code: http.StatusUnauthorized})
		return
	}

	c.JSON(http.StatusOK, models.GetAllNotesResponse{Notes: notes})
}

func (s *NotesController) AddNoteHandler(c *gin.Context) {
	req := models.AddNoteRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "invalid request format", Code: http.StatusBadRequest})
		return
	}

	id, err := s.NotesService.AddNote(req.SID.SID, req.Note)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: err.Error(), Code: http.StatusUnauthorized})
		return
	}

	c.JSON(http.StatusOK, models.AddNoteResponse{ID: id})
}

func (s *NotesController) DeleteNoteHandler(c *gin.Context) {
	req := models.DeleteNoteRequest{}

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "invalid request format", Code: http.StatusBadRequest})
		return
	}

	err = s.NotesService.DeleteNote(req.SID.SID, req.ID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: err.Error(), Code: http.StatusUnauthorized})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{Message: "success"})
}
