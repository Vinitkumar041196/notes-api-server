package service

import (
	"notes-api-server/internal/app"
	"notes-api-server/internal/models"
	"notes-api-server/internal/store"
)

type NotesService struct {
	NotesStore   *store.NotesStore
	SessionStore *store.SessionStore
}

func NewNotesService(app *app.App) models.NotesService {
	return &NotesService{
		NotesStore:   app.NotesStore,
		SessionStore: app.SessionStore,
	}
}

func (n *NotesService) AddNote(sid, note string) (uint32, error) {
	user, err := n.SessionStore.GetSession(sid)
	if err != nil {
		return 0, err
	}

	return n.NotesStore.AddNote(user, note), nil
}

func (n *NotesService) GetAllNotes(sid string) ([]models.Note, error) {
	user, err := n.SessionStore.GetSession(sid)
	if err != nil {
		return nil, err
	}

	return n.NotesStore.GetAllNotes(user), nil
}

func (n *NotesService) DeleteNote(sid string, id uint32) error {
	user, err := n.SessionStore.GetSession(sid)
	if err != nil {
		return err
	}

	n.NotesStore.DeleteNote(user, id)
	return nil
}
