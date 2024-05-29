package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/jetnoli/notion-voice-assistant/config/client"
	"github.com/jetnoli/notion-voice-assistant/wrappers"
)

func ServeRoot(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	res, err := client.WhisperApi.Get("/", wrappers.ApiGetRequestOptions{})

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(502)
		return
	}

	defer res.Body.Close()

	absPath, err := filepath.Abs("static/html/index.html")

	if err != nil {
		http.Error(w, "Error Getting Path To file:\n"+err.Error(), http.StatusInternalServerError)
		return
	}

	html, err := os.ReadFile(absPath)

	if err != nil {
		http.Error(w, "Error Reading file:\n"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(html))
}
