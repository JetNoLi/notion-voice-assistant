package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jetnoli/notion-voice-assistant/config"
	"github.com/jetnoli/notion-voice-assistant/services"
)

type NotionDatabase struct {
	Object  string `json:"object"`
	Results any    `json:"name"`
}

func GetDatabases(w http.ResponseWriter, r *http.Request) {

	data, err := services.GetDatabases[any]()

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("Content-Type", "application/json")
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

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)

}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	databaseId := config.NotionMainDbId

	data, err := services.CreateDatabaseItem[any](databaseId, &services.ItemData{Title: "Planning 101", Status: "Planning"})

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
