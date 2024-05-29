package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jetnoli/notion-voice-assistant/services"
)

type AssistRequestBody struct {
	Prompt string `json:"prompt"`
}

func Assist(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var reqBody AssistRequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf(`{
			"title": "Error: Bad Request",
			"message": "%s"
		}`, err.Error())))
	}

	data, err := services.Assist(reqBody.Prompt)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf(`{
			"title": "Error: Bad Request",
			"message": "%s"
		}`, err.Error())))
	}

	w.Header().Add("Content-Type", "application/json")

	w.Write([]byte(data))
}
