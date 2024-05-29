package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/jetnoli/notion-voice-assistant/config"
	"github.com/jetnoli/notion-voice-assistant/handlers"
	routes "github.com/jetnoli/notion-voice-assistant/routes"
	"github.com/jetnoli/notion-voice-assistant/utils"
)

func main() {
	isPortDefined, err := regexp.MatchString("^[0-9]{1,45}$", config.Port)

	if err != nil {
		fmt.Println("Port Env Variable Cannot Be Parsed")
		panic(err.Error())
	}

	utils.Assert(isPortDefined, "Port is not defined")

	router := http.NewServeMux()

	router.Handle("/notion/", routes.NotionRouter())
	router.Handle("/completion/", routes.GptRouter())
	router.Handle("/transcribe/", routes.WhisperRouter())
	router.Handle("/", routes.HTMLRouter())
	router.HandleFunc("/health/{$}", handlers.HealthCheck)

	server := http.Server{
		Addr:         ":" + config.Port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      router,
	}

	fmt.Println("Starting Server on http://localhost:" + config.Port)

	log.Fatal(server.ListenAndServe())
}
