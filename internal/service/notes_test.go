package service

import (
	"notes-api-server/internal/app"
	"notes-api-server/internal/models"
	"reflect"
	"testing"
)

func TestNotesService_AddNote(t *testing.T) {
	type args struct {
		sid  string
		note string
	}
	app := app.NewApp()
	newSID := app.SessionStore.AddSession("test@example.com")

	notesServ := NewNotesService(app)

	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		{
			name: "Valid Add Note",
			args: args{
				sid:  newSID,
				note: "this is first note",
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Invalid Session",
			args: args{
				sid:  "b6818389-9371-45a5-b9c6-09c91e92a487",
				note: "this is first note",
			},
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := notesServ.AddNote(tt.args.sid, tt.args.note)
			if (err != nil) != tt.wantErr {
				t.Errorf("NotesService.AddNote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NotesService.AddNote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotesService_GetAllNotes(t *testing.T) {
	type args struct {
		sid string
	}

	app := app.NewApp()
	newSID := app.SessionStore.AddSession("test@example.com")

	notesServ := NewNotesService(app)
	id, _ := notesServ.AddNote(newSID, "test note")

	tests := []struct {
		name    string
		args    args
		want    []models.Note
		wantErr bool
	}{
		{
			name: "Valid Get All Notes",
			args: args{
				sid: newSID,
			},
			want: []models.Note{
				{
					ID:   id,
					Note: "test note",
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid Session",
			args: args{
				sid: "b6818389-9371-45a5-b9c6-09c91e92a487",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := notesServ.GetAllNotes(tt.args.sid)
			if (err != nil) != tt.wantErr {
				t.Errorf("NotesService.GetAllNotes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NotesService.GetAllNotes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotesService_DeleteNote(t *testing.T) {
	type args struct {
		sid string
		id  uint32
	}

	app := app.NewApp()
	newSID := app.SessionStore.AddSession("test@example.com")

	notesServ := NewNotesService(app)
	id, _ := notesServ.AddNote(newSID, "test note")

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid Delete Note",
			args: args{
				sid: newSID,
				id:  id,
			},
			wantErr: false,
		},
		{
			name: "Invalid Session",
			args: args{
				sid: "b6818389-9371-45a5-b9c6-09c91e92a487",
				id:  1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := notesServ.DeleteNote(tt.args.sid, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("NotesService.DeleteNote() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
