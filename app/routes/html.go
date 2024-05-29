package routes

import (
	"net/http"

	html "github.com/jetnoli/notion-voice-assistant/handlers"
	"github.com/jetnoli/notion-voice-assistant/wrappers"
)

func HTMLRouter() *http.ServeMux {
	router := wrappers.CreateRouter("/", wrappers.RouterOptions{})

	router.Get("/", html.ServeRoot)

	return router.Mux
}
