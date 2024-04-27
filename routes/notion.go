package routes

import (
	"net/http"

	"github.com/jetnoli/notion-voice-assistant/handlers"
	notion "github.com/jetnoli/notion-voice-assistant/handlers"
	"github.com/jetnoli/notion-voice-assistant/wrappers"
)

func NotionRouter() *http.ServeMux {

	router := wrappers.CreateRouter("/notion", wrappers.RouterOptions{ExactPathsOnly: true})

	router.Get("/", notion.GetDatabases)
	router.Get("/{id}", notion.GetDatabaseById)
	router.Post("/", handlers.CreateTask)

	return router.Mux
}
