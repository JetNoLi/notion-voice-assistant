package routes

import (
	"net/http"

	"github.com/jetnoli/notion-voice-assistant/handlers"
	"github.com/jetnoli/notion-voice-assistant/wrappers"
)

func UserRouter() *http.ServeMux {
	router := wrappers.CreateRouter("/user", wrappers.RouterOptions{
		ExactPathsOnly: false,
	})

	router.Post("/signup", handlers.SignUp)
	router.Get("/", handlers.GetAllUsers)
	router.Get("/{id}", handlers.GetUserById)
	router.Get("/username/{username}", handlers.GetUserByUsername)
	router.Put("/{id}", handlers.UpdateUserById)
	router.Delete("/{id}", handlers.DeleteUserById)
	router.Delete("/", handlers.DeleteAllUsers)

	return router.Mux
}