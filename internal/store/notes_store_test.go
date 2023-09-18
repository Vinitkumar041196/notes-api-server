package store

import (
	"notes-api-server/internal/models"
	"reflect"
	"testing"
)

func TestNotesStore_GetNextNoteID(t *testing.T) {
	n := NewNotesStore()

	tests := []struct {
		name string
		want uint32
	}{
		{
			name: "Valid",
			want: 1,
		},
		{
			name: "Next Valid",
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := n.GetNextNoteID(); got != tt.want {
				t.Errorf("NotesStore.GetNextNoteID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotesStore_AddNote(t *testing.T) {
	type args struct {
		user string
		note string
	}

	n := NewNotesStore()

	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "Valid Add Note",
			args: args{
				user: "test@example.com",
				note: "test note",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := n.AddNote(tt.args.user, tt.args.note); got != tt.want {
				t.Errorf("NotesStore.AddNote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotesStore_GetAllNotes(t *testing.T) {
	type args struct {
		user string
	}

	n := NewNotesStore()
	id := n.AddNote("test@example.com", "test note")

	tests := []struct {
		name string
		args args
		want []models.Note
	}{
		{
			name: "Valid Get All Notes",
			args: args{
				user: "test@example.com",
			},
			want: []models.Note{
				{
					ID:   id,
					Note: "test note",
				},
			},
		},
		{
			name: "User with no notes",
			args: args{
				user: "test2@example.com",
			},
			want: []models.Note{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := n.GetAllNotes(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NotesStore.GetAllNotes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotesStore_DeleteNote(t *testing.T) {
	type args struct {
		user string
		id   uint32
	}

	n := NewNotesStore()
	id := n.AddNote("test@example.com", "test note")

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Valid",
			args: args{
				user: "test@example.com",
				id:   id,
			},
		},
		{
			name: "Non existing user",
			args: args{
				user: "test123@example.com",
				id:   id,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n.DeleteNote(tt.args.user, tt.args.id)
		})
	}
}
