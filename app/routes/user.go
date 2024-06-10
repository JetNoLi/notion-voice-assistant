package routes

import (
	"net/http"

	"github.com/jetnoli/notion-voice-assistant/handlers"
	Router "github.com/jetnoli/notion-voice-assistant/wrappers/router"
)

func UserRouter() *http.ServeMux {
	router := Router.CreateRouter("/user", Router.RouterOptions{
		ExactPathsOnly: false,
	})

	router.Post("/signup", handlers.SignUp, &Router.RouteOptions{})
	router.Get("/", handlers.GetAllUsers, &Router.RouteOptions{})
	router.Get("/{id}", handlers.GetUserById, &Router.RouteOptions{})
	router.Get("/username/{username}", handlers.GetUserByUsername, &Router.RouteOptions{})
	router.Put("/{id}", handlers.UpdateUserById, &Router.RouteOptions{})
	router.Delete("/{id}", handlers.DeleteUserById, &Router.RouteOptions{})
	router.Delete("/", handlers.DeleteAllUsers, &Router.RouteOptions{})

	return router.Mux
}
