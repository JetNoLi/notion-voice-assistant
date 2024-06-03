package router

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/jetnoli/notion-voice-assistant/wrappers/serve"
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

// Templating

// Serve html at the given filepath relative to app
func (router Router) Serve(path string, filePath string) {
	route := router.CreatePath(path, "GET")

	router.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		html, err := serve.Html(filePath)

		if err != nil {
			http.Error(w, "Error Reading file:\n"+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(html)
	})
}

// Serves all html in given directory relative to app
// Ignores nested directories
// Include trailing slash in dir name
func (router Router) ServeDir(baseUrlPath string, dirPath string) {
	absPath, err := filepath.Abs(dirPath)

	if err != nil {
		panic("error reading dir: " + err.Error())
	}

	files, err := os.ReadDir(absPath)

	if err != nil {
		panic("failed to read dir for html pages: " + err.Error())
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()
		filePath := absPath + "/" + fileName

		route := baseUrlPath + strings.Split(fileName, ".")[0] + "/"

		if route == "/index/" {
			route = "/"
		}

		router.Serve(route, filePath)
	}
}
