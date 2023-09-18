package server

import (
	"net/http"
	"notes-api-server/internal/app"
	"notes-api-server/internal/controllers"
	"notes-api-server/internal/service"

	"github.com/gin-gonic/gin"
)

func NewRouter(app *app.App) *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	NewSignupRouter(app, router)
	NewLoginRouter(app, router)
	NewNotesRouter(app, router)

	return router
}

func NewSignupRouter(app *app.App, router *gin.Engine) {
	signupController := controllers.SignupController{
		UserService: service.NewUserService(app),
	}

	router.POST("/signup", signupController.SignupHandler)
}

func NewLoginRouter(app *app.App, router *gin.Engine) {
	loginController := controllers.LoginController{
		UserService: service.NewUserService(app),
	}

	router.POST("/login", loginController.LoginHandler)
}

func NewNotesRouter(app *app.App, router *gin.Engine) {
	notesController := controllers.NotesController{
		NotesService: service.NewNotesService(app),
	}

	router.GET("/notes", notesController.GetAllNotesHandler)
	router.POST("/notes", notesController.AddNoteHandler)
	router.DELETE("/notes", notesController.DeleteNoteHandler)

}
