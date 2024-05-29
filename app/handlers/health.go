package handlers

import (
	"fmt"
	"io"
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

	data, err := io.ReadAll(res.Body)

	if err != nil {
		w.WriteHeader(500)
	}

	w.Write(data)
}
