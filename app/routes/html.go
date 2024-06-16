package routes

import (
	"net/http"

	html "github.com/jetnoli/notion-voice-assistant/handlers"
	"github.com/jetnoli/notion-voice-assistant/middleware"

	home "github.com/jetnoli/notion-voice-assistant/view/pages/home"
	login "github.com/jetnoli/notion-voice-assistant/view/pages/login"
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

	router.ServeTempl(map[string]*Router.TemplPage{
		"/": {
			PageComponent: home.Index(),
			Options: &Router.RouteOptions{
				PreHandlerMiddleware: []Router.MiddlewareHandler{middleware.CheckAuthorization},
			},
		},
		"/login": {
			PageComponent: login.Login(),
		},
		"/signup": {
			PageComponent: signup.Signup(),
		},
	})

	router.Post("/htmx/signup", html.SignUpHtmx, &Router.RouteOptions{})
	router.Post("/htmx/signin", html.SignInHtmx, &Router.RouteOptions{})

	return router.Mux
}
