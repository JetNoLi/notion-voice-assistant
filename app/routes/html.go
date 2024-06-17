package routes

import (
	"net/http"

	"github.com/jetnoli/notion-voice-assistant/handlers"
	"github.com/jetnoli/notion-voice-assistant/middleware"

	"github.com/jetnoli/notion-voice-assistant/view/pages/home"
	"github.com/jetnoli/notion-voice-assistant/view/pages/login"
	"github.com/jetnoli/notion-voice-assistant/view/pages/signup"

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

	router.ServeTempl(map[string]*Router.TemplPage{
		"/": {
			PageComponent: home.Page(),
			Options: &Router.RouteOptions{
				PreHandlerMiddleware: []Router.MiddlewareHandler{middleware.CheckAuthorization},
			},
		},
		"/login": {
			PageComponent: login.Page(),
		},
		"/signup": {
			PageComponent: signup.Page(),
		},
	})

	router.Post("/htmx/signup", handlers.SignUpHtmx, &Router.RouteOptions{})
	router.Post("/htmx/signin", handlers.SignInHtmx, &Router.RouteOptions{})

	return router.Mux
}
