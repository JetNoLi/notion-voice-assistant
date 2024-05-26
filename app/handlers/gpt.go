package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jetnoli/notion-voice-assistant/services"
)

type AssistRequestBody struct {
	prompt string
}

func Assist(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering assist")
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

	res, err := services.Assist(reqBody.prompt)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf(`{
			"title": "Error: Bad Request",
			"message": "%s"
		}`, err.Error())))
	}

	w.Write([]byte(fmt.Sprintf(`{
		"response": "%s"
	}`, res)))
}
