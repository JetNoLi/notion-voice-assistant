package routes

import (
	"net/http"

	gpt "github.com/jetnoli/notion-voice-assistant/handlers"
	"github.com/jetnoli/notion-voice-assistant/wrappers/router"
)

func GptRouter() *http.ServeMux {
	router := router.CreateRouter("/completion", router.RouterOptions{})

	router.Post("/", gpt.Assist)

	return router.Mux
}
