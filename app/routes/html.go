package routes

import (
	"net/http"

	"github.com/a-h/templ"
	html "github.com/jetnoli/notion-voice-assistant/handlers"

	home "github.com/jetnoli/notion-voice-assistant/view/pages/home"
	signup "github.com/jetnoli/notion-voice-assistant/view/pages/signup"

	Router "github.com/jetnoli/notion-voice-assistant/wrappers/router"
)

func HTMLRouter() *http.ServeMux {
	router := Router.CreateRouter("/", Router.RouterOptions{
		ExactPathsOnly: true,
	})

	// Serve Styles for Pages
	router.ServeDir("/", "view/pages/", &Router.ServeDirOptions{
		IncludedExtensions:         []string{".css"},
		Recursive:                  true,
		RoutePathContainsExtension: true,
	})

	// Serve Styles for Components
	router.ServeDir("/", "view/components/", &Router.ServeDirOptions{
		IncludedExtensions:         []string{".css"},
		Recursive:                  true,
		RoutePathContainsExtension: true,
	})

	router.ServeDir("/assets/", "assets/", &Router.ServeDirOptions{
		Recursive:                  true,
		RoutePathContainsExtension: true,
	})

	//TODO: Make Templ Handler
	router.Mux.Handle("/", templ.Handler(home.Index()))
	router.Mux.Handle("/signup/", templ.Handler(signup.Signup()))
	router.Post("/htmx/signup", html.SignUpHtmx, &Router.RouteOptions{})
	router.Post("/htmx/signin", html.SignInHtmx, &Router.RouteOptions{})

	return router.Mux
}
