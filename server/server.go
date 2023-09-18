package server

import (
	"crypto/tls"
	"log"
	"net/http"
	"notes-api-server/internal/app"
)

func Init(app *app.App) {
	router := NewRouter(app)

	server := &http.Server{
		Addr:    app.Env.ServerAddress,
		Handler: router,
	}

	log.Printf("Starting server at: %s", server.Addr)

	if app.Env.Config.GetBool("ENABLE_TLS") {
		log.Fatal(server.ListenAndServeTLS(app.Env.Config.GetString("SSL_CRT_PATH"), app.Env.Config.GetString("SSL_KEY_PATH")))
	} else {
		//Globally disabling SSL certificate check
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		log.Fatal(server.ListenAndServe())
	}
}
