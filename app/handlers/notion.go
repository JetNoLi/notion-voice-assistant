package handlers

import (
	"encoding/json"
	"fmt"
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
	prompt, err := TranscribeResponse(r)

	if err != nil {

		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	databaseId := config.NotionMainDbId

	fmt.Println("fetching related content")
	relatedContent, err := services.GetAllRelatedPages(databaseId)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Println("creating page from related content")

	data, err := services.CreatePageFromRelatedContent(databaseId, relatedContent, prompt.Result)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(&data)

	if err != nil {
		http.Error(w, "Error Returning Data: "+err.Error(), http.StatusInternalServerError)
		return
	}

}
