package routes

import (
	"net/http"

	"github.com/jetnoli/notion-voice-assistant/handlers"
	"github.com/jetnoli/notion-voice-assistant/wrappers"
)

func WhisperRouter() *http.ServeMux {

	router := wrappers.CreateRouter("/transcribe", wrappers.RouterOptions{ExactPathsOnly: true})

	router.Post("/", handlers.Transcribe)

	return router.Mux
}
