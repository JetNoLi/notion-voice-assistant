package handlers

import (
	"fmt"
	"net/http"

	"github.com/jetnoli/notion-voice-assistant/config/client"
	"github.com/jetnoli/notion-voice-assistant/wrappers"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	res, err := client.WhisperApi.Get("/", wrappers.ApiGetRequestOptions{})

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(502)
		return
	}

	defer res.Body.Close()

	w.Write([]byte(`<span id="health-check-indicator">&#10003;</span>`))
}
