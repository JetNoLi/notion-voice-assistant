package routes

import (
	"net/http"

	html "github.com/jetnoli/notion-voice-assistant/handlers"
	"github.com/jetnoli/notion-voice-assistant/wrappers/router"
)

func HTMLRouter() *http.ServeMux {
	router := router.CreateRouter("/", router.RouterOptions{
		ExactPathsOnly: true,
	})

	router.ServeDir("/", "static/html/pages/")

	router.Post("/htmx/signup", html.SignupHtmx)

	return router.Mux
}
