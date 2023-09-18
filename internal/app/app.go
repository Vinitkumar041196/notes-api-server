package app

import (
	"log"
	"notes-api-server/internal/store"

	"github.com/spf13/viper"
)

type App struct {
	Env          *Env
	UserStore    *store.UserStore
	NotesStore   *store.NotesStore
	SessionStore *store.SessionStore
}

func NewApp() *App {
	app := new(App)
	app.Env = NewEnv()
	app.UserStore = store.NewUserStore()
	app.NotesStore = store.NewNotesStore()
	app.SessionStore = store.NewSessionStore()
	return app
}

type Env struct {
	AppEnv        string       `mapstructure:"APP_ENV"`
	ServerAddress string       `mapstructure:"SERVER_ADDRESS"`
	Config        *viper.Viper `mapstructure:"-"`
}

func NewEnv() *Env {
	conf := viper.New()

	conf.AutomaticEnv()

	conf.BindEnv("APP_ENV")
	conf.BindEnv("SERVER_ADDRESS")
	conf.BindEnv("ACCESS_TOKEN_EXPIRY_MIN")
	conf.BindEnv("REFRESH_TOKEN_EXPIRY_MIN")
	conf.BindEnv("ACCESS_TOKEN_SECRET")
	conf.BindEnv("REFRESH_TOKEN_SECRET")
	conf.BindEnv("JWT_SIGNING_METHOD")

	env := Env{}
	err := conf.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}
	env.Config = conf

	if env.AppEnv == "development" {
		log.Println("App is running in development env")
	}

	return &env
}
