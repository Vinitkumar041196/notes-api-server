package service

import (
	"notes-api-server/internal/app"
	"notes-api-server/internal/models"
	"reflect"
	"testing"
)

func TestUserService_AddUser(t *testing.T) {
	type args struct {
		u models.User
	}

	app := app.NewApp()
	userServ := NewUserService(app)

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid Add User",
			args: args{
				u: models.User{
					Name:     "user",
					Email:    "user@example.com",
					Password: "pass@123",
				},
			},
			wantErr: false,
		},
		{
			name: "Add Existing User",
			args: args{
				u: models.User{
					Name:     "user",
					Email:    "user@example.com",
					Password: "pass@123",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := userServ.AddUser(tt.args.u); (err != nil) != tt.wantErr {
				t.Errorf("UserService.AddUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserService_GetUser(t *testing.T) {
	type args struct {
		email string
	}

	app := app.NewApp()
	user:=models.User{
		Name:     "user",
		Email:    "user@example.com",
		Password: "pass@123",
	}
	app.UserStore.AddUser(user)

	userServ := NewUserService(app)

	tests := []struct {
		name    string
		args    args
		want    models.User
		wantErr bool
	}{
		{
			name: "Valid Get User",
			args: args{
				email: "user@example.com",
			},
			want: user,
			wantErr: false,
		},
		{
			name: "Non existing user",
			args: args{
				email: "user123@example.com",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := userServ.GetUser(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_CreateUserSession(t *testing.T) {
	type args struct {
		email string
	}

	app := app.NewApp()
	userServ := NewUserService(app)

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Valid User Session",
			args: args{
				email: "user@example.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := userServ.CreateUserSession(tt.args.email); len(got) == 0 {
				t.Errorf("UserService.CreateUserSession() = %v, want non empty uuid", got)
			}
		})
	}
}
