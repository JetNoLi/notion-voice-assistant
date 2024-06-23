package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/jetnoli/notion-voice-assistant/config"
	"github.com/jetnoli/notion-voice-assistant/db"
	"github.com/jetnoli/notion-voice-assistant/handlers"
	"github.com/jetnoli/notion-voice-assistant/middleware"
	"github.com/jetnoli/notion-voice-assistant/routes"
	"github.com/jetnoli/notion-voice-assistant/utils"
	Router "github.com/jetnoli/notion-voice-assistant/wrappers/router"
)

func main() {

	isPortDefined, err := regexp.MatchString("^[0-9]{1,45}$", config.Port)

	if err != nil {
		fmt.Println("Port Env Variable Cannot Be Parsed")
		panic(err.Error())
	}

	utils.Assert(isPortDefined, "Port is not defined")

	db.Connect()
	defer db.Db.Close()

	router := Router.CreateRouter("*", Router.RouterOptions{
		PreHandlerMiddleware: []Router.MiddlewareHandler{middleware.DecodeToken},
	})

	router.Handle("/notion/", routes.NotionRouter())
	router.Handle("/completion/", routes.GptRouter())
	router.Handle("/transcribe/", routes.WhisperRouter())
	router.Handle("/user/", routes.UserRouter())
	router.Handle("/", routes.HTMLRouter())
	router.HandleFunc("/health/{$}", handlers.HealthCheck, &Router.RouteOptions{})

	server := http.Server{
		Addr:         ":" + config.Port,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Handler:      router.Mux,
	}

	fmt.Println("Starting Server on http://localhost:" + config.Port)

	log.Fatal(server.ListenAndServe())
}
