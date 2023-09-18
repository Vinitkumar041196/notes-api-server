package store

import (
	"notes-api-server/internal/models"
	"sync"
)

type NotesStore struct {
	noteID uint32
	store  map[string]map[uint32]models.Note
	mux    sync.Mutex
}

func NewNotesStore() *NotesStore {
	n := new(NotesStore)
	n.store = make(map[string]map[uint32]models.Note)
	n.noteID = 0
	return n
}

func (n *NotesStore) GetNextNoteID() uint32 {
	n.noteID += 1
	return n.noteID
}

func (n *NotesStore) AddNote(user, note string) uint32 {
	n.mux.Lock()
	id := n.GetNextNoteID()
	n.mux.Unlock()

	notes := n.store[user]
	if notes == nil {
		notes = make(map[uint32]models.Note)
	}
	notes[id] = models.Note{ID: id, Note: note}

	n.mux.Lock()
	n.store[user] = notes
	n.mux.Unlock()

	return id
}

func (n *NotesStore) GetAllNotes(user string) []models.Note {
	notes := n.store[user]
	notesArr := []models.Note{}
	for _, n := range notes {
		notesArr = append(notesArr, n)
	}
	return notesArr
}

func (n *NotesStore) DeleteNote(user string, id uint32) {
	notes := n.store[user]
	delete(notes, id)

	n.mux.Lock()
	n.store[user] = notes
	n.mux.Unlock()
}
