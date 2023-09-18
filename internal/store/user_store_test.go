package store

import (
	"notes-api-server/internal/models"
	"reflect"
	"testing"
)

func TestUserStore_AddUser(t *testing.T) {
	type args struct {
		user models.User
	}

	u := NewUserStore()

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid Add User",
			args: args{
				user: models.User{
					Name:     "user",
					Email:    "test@example.com",
					Password: "pass@123",
				},
			},
			wantErr: false,
		},
		{
			name: "Add Existing User",
			args: args{
				user: models.User{
					Name:     "user",
					Email:    "test@example.com",
					Password: "pass@123",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := u.AddUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("UserStore.AddUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserStore_GetUser(t *testing.T) {
	type args struct {
		email string
	}

	u := NewUserStore()
	user := models.User{
		Name:     "user",
		Email:    "test@example.com",
		Password: "pass@123",
	}
	u.AddUser(user)

	tests := []struct {
		name    string
		args    args
		want    models.User
		wantErr bool
	}{
		{
			name: "Valid Get User",
			args: args{
				email: "test@example.com",
			},
			want:    user,
			wantErr: false,
		},
		{
			name: "Get Non Existing1 User",
			args: args{
				email: "test123@example.com",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := u.GetUser(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserStore.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserStore.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
