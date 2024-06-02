package routes

import (
	"net/http"

	"github.com/jetnoli/notion-voice-assistant/handlers"
	"github.com/jetnoli/notion-voice-assistant/wrappers/router"
)

func WhisperRouter() *http.ServeMux {

	router := router.CreateRouter("/transcribe", router.RouterOptions{ExactPathsOnly: true})

	router.Post("/", handlers.Transcribe)

	return router.Mux
}
