package routes

import (
	"net/http"

	notion "github.com/jetnoli/notion-voice-assistant/handlers"
	Router "github.com/jetnoli/notion-voice-assistant/wrappers/router"
)

func NotionRouter() *http.ServeMux {

	router := Router.CreateRouter("/notion", Router.RouterOptions{ExactPathsOnly: true})

	router.Get("/", notion.GetDatabases, &Router.RouteOptions{})
	router.Get("/{id}", notion.GetDatabaseById, &Router.RouteOptions{})
	router.Post("/", notion.CreateTask, &Router.RouteOptions{})

	return router.Mux
}
