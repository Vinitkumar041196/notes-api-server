package models

type Note struct {
	ID   uint32 `json:"id"`
	Note string `json:"note"`
}

type NotesService interface {
	AddNote(sid, note string) (uint32, error)
	GetAllNotes(sid string) ([]Note, error)
	DeleteNote(sid string, id uint32) error
}
