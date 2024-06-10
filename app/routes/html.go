package routes

import (
	"net/http"

	html "github.com/jetnoli/notion-voice-assistant/handlers"
	Router "github.com/jetnoli/notion-voice-assistant/wrappers/router"
)

func HTMLRouter() *http.ServeMux {
	router := Router.CreateRouter("/", Router.RouterOptions{
		ExactPathsOnly: true,
	})

	router.ServeDir("/", "static/html/pages/")

	router.Post("/htmx/signup", html.SignUpHtmx, &Router.RouteOptions{})
	router.Post("/htmx/signin", html.SignInHtmx, &Router.RouteOptions{})

	return router.Mux
}
