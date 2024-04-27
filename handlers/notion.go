package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jetnoli/notion-voice-assistant/wrappers"
)

// TODO: Type Properly
type NotionDatabase struct {
	Object  string `json:object`
	Results any    `json:name`
}

var notionApi = wrappers.Api{
	BaseUrl: "https://api.notion.com/v1",
	Client:  &http.Client{},
	Headers: map[string]string{
		"Content-Type":   "application/json",
		"Authorization":  "",
		"Notion-Version": "2022-06-28",
	},
}

func GetDatabases(w http.ResponseWriter, r *http.Request) {
	body := []byte(`{
    	"filter": {
     	   "value": "database",
      	  "property": "object"
    	},
    	"sort": {
    	    "direction": "ascending",
    	    "timestamp": "last_edited_time"
    	}
	}`)

	res, err := notionApi.Post("/search", body, wrappers.ApiPostRequestOptions{})

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	defer res.Body.Close()

	json.NewEncoder(w).Encode(res.Body)
}

func GetDatabaseById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	res, err := notionApi.Get("/databases/"+id, wrappers.ApiGetRequestOptions{})

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	defer res.Body.Close()

	json.NewEncoder(w).Encode(res.Body)

}

func CreateTask(w http.ResponseWriter, r *http.Request) {

}
