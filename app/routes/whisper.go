package routes

import (
	"net/http"

	"github.com/jetnoli/notion-voice-assistant/handlers"
	Router "github.com/jetnoli/notion-voice-assistant/wrappers/router"
)

func WhisperRouter() *http.ServeMux {

	router := Router.CreateRouter("/transcribe", Router.RouterOptions{ExactPathsOnly: true})

	router.Post("/", handlers.Transcribe, &Router.RouteOptions{})

	return router.Mux
}
