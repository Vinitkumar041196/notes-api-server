package store

import (
	"testing"
)

func TestSessionStore_AddSession(t *testing.T) {
	type args struct {
		email string
	}

	s := NewSessionStore()

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Valid Add Session",
			args: args{
				email: "test@example.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s.AddSession(tt.args.email); len(got) == 0 {
				t.Errorf("SessionStore.AddSession() = %v, want non empty uuid", got)
			}
		})
	}
}

func TestSessionStore_GetSession(t *testing.T) {
	type args struct {
		sid string
	}

	s := NewSessionStore()
	sid := s.AddSession("test@example.com")
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Valid Get Session",
			args: args{
				sid: sid,
			},
			want: "test@example.com",
			wantErr: false,
		},
		{
			name: "Invalid Session ID",
			args: args{
				sid: "12324-324-3243-43234",
			},
			want: "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetSession(tt.args.sid)
			if (err != nil) != tt.wantErr {
				t.Errorf("SessionStore.GetSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SessionStore.GetSession() = %v, want %v", got, tt.want)
			}
		})
	}
}
