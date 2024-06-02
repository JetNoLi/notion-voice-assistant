package routes

import (
	"net/http"

	notion "github.com/jetnoli/notion-voice-assistant/handlers"
	"github.com/jetnoli/notion-voice-assistant/wrappers/router"
)

func NotionRouter() *http.ServeMux {

	router := router.CreateRouter("/notion", router.RouterOptions{ExactPathsOnly: true})

	router.Get("/", notion.GetDatabases)
	router.Get("/{id}", notion.GetDatabaseById)
	router.Post("/", notion.CreateTask)

	return router.Mux
}
