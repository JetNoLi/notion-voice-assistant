package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jetnoli/notion-voice-assistant/handlers"
	routes "github.com/jetnoli/notion-voice-assistant/routes"
)

var port string = "8080"

func main() {
	router := http.NewServeMux()

	router.Handle("/notion/", routes.NotionRouter())
	router.Handle("/completion/", routes.GptRouter())
	router.HandleFunc("GET /{$}", handlers.HealthCheck)

	server := http.Server{
		Addr:         ":" + port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      router,
	}

	fmt.Println("Starting Server on http://localhost:" + port)

	log.Fatal(server.ListenAndServe())
}
