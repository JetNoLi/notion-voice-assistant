package wrappers

import (
	"fmt"
	"net/http"
)

type RouterOptions struct {
	ExactPathsOnly bool // Appends the {$} for all paths in router
}

type Router struct {
	Path    string
	Mux     *http.ServeMux
	Options RouterOptions
}

// TODO: Add Global Response Headers
func CreateRouter(path string, options RouterOptions) *Router {
	router := &Router{}
	router.Mux = http.NewServeMux()
	router.Path = path
	router.Options = options

	return router
}

func (router Router) CreatePath(path string, method string) string {
	pathEndString := ""

	if router.Options.ExactPathsOnly {
		end := len(path) - 1

		if len(path) > 3 && path[end-3:end] == "{$}" {
			pathEndString = ""
		} else if path[len(path)-1] == '/' {
			pathEndString = "{$}"
		} else {
			pathEndString = "/{$}"
		}
	}

	url := method + " " + router.Path

	// To avoid double // in request e.g. GET //path-name
	if router.Path[len(router.Path)-1] == '/' && path[0] == '/' {
		url += path[1:]
	} else {
		url += path
	}

	fmt.Println("Registering: " + url + pathEndString)

	return url + pathEndString
}

// Adds Basic Logging to handler
func (router Router) HandleFunc(path string, handler http.HandlerFunc) {
	router.Mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Starting Request: " + path)
		handler(w, r)
		fmt.Println("Completing Request: " + path)
	})
}

func (router Router) Get(path string, handler http.HandlerFunc) {
	route := router.CreatePath(path, "GET")
	router.HandleFunc(route, handler)
}

func (router Router) Post(path string, handler http.HandlerFunc) {
	route := router.CreatePath(path, "POST")
	router.HandleFunc(route, handler)
}

func (router Router) Delete(path string, handler http.HandlerFunc) {
	route := router.CreatePath(path, "DELETE")
	router.HandleFunc(route, handler)
}

func (router Router) Put(path string, handler http.HandlerFunc) {
	route := router.CreatePath(path, "PUT")
	router.HandleFunc(route, handler)
}

func (router Router) Patch(path string, handler http.HandlerFunc) {
	route := router.CreatePath(path, "PATCH")
	router.HandleFunc(route, handler)
}
