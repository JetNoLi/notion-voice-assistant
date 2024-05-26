package routes

import (
	"net/http"

	gpt "github.com/jetnoli/notion-voice-assistant/handlers"
	"github.com/jetnoli/notion-voice-assistant/wrappers"
)

func GptRouter() *http.ServeMux {
	router := wrappers.CreateRouter("/completion", wrappers.RouterOptions{})

	router.Post("/", gpt.Assist)

	return router.Mux
}
