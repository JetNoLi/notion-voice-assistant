package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jetnoli/notion-voice-assistant/services"
)

type NotionDatabase struct {
	Object  string `json:object`
	Results any    `json:name`
}

func GetDatabases(w http.ResponseWriter, r *http.Request) {

	data, err := services.GetDatabases[any]()

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(data)
}

func GetDatabaseById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	data, err := services.GetDatabaseById[any](id)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(data)

}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	// Connect to Correct Notion Workspace
	// Create Struct to Add Task
}
