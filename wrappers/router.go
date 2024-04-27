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

	return method + " " + router.Path + path + pathEndString
}

func (router Router) Get(path string, handler http.HandlerFunc) {
	fmt.Println("Registering Route: ")
	route := router.CreatePath(path, "GET")
	fmt.Println(route)
	router.Mux.HandleFunc(route, handler)
}

func (router Router) Post(path string, handler http.HandlerFunc) {
	route := router.CreatePath(path, "POST")
	router.Mux.HandleFunc(route, handler)
}

func (router Router) Delete(path string, handler http.HandlerFunc) {
	route := router.CreatePath(path, "DELETE")
	router.Mux.HandleFunc(route, handler)
}

func (router Router) Put(path string, handler http.HandlerFunc) {
	route := router.CreatePath(path, "PUT")
	router.Mux.HandleFunc(route, handler)
}

func (router Router) Patch(path string, handler http.HandlerFunc) {
	route := router.CreatePath(path, "PATCH")
	router.Mux.HandleFunc(route, handler)
}
