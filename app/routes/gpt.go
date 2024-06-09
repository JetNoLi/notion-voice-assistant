package routes

import (
	"net/http"

	"github.com/jetnoli/notion-voice-assistant/handlers"
	Router "github.com/jetnoli/notion-voice-assistant/wrappers/router"
)

func GptRouter() *http.ServeMux {
	router := Router.CreateRouter("/completion", Router.RouterOptions{})

	router.Post("/", handlers.Assist, &Router.RouteOptions{})

	return router.Mux
}
