package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jetnoli/notion-voice-assistant/config"
	"github.com/jetnoli/notion-voice-assistant/wrappers"
)

var notionApi = wrappers.Api{
	BaseUrl: "https://api.notion.com/v1",
	Client:  &http.Client{},
	Headers: map[string]string{
		"Content-Type":   "application/json",
		"Authorization":  "Bearer " + config.NotionApiKey,
		"Notion-Version": "2022-06-28",
	},
}

func GetDatabases[T any]() (data T, err error) {
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
		return data, err
	}

	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&data)
	return data, err
}

func GetDatabaseById[T any](id string) (data T, err error) {
	res, err := notionApi.Get("/databases/"+id, wrappers.ApiGetRequestOptions{})

	if err != nil {
		return data, err
	}

	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&data)

	return data, err
}

type ItemData struct {
	Title  string
	Status string
}

func CreateDatabaseItem[R any](databaseId string, itemData *ItemData) (item *R, err error) {
	reqStr := fmt.Sprintf(`{
		"parent": {
			"database_id": "%s"
		},
		"icon": {
			"emoji": "🥬"
		},
		"properties": {
			"Name": {
				"title": [
					{
						"text": {
							"content": "%s"
						}
					}
				]
			},
			"Tags": {
				"multi_select": [
					{
						"name": "Workout"
					}
				]
			}

		}
	}`, databaseId, itemData.Title)

	fmt.Println(reqStr)

	body := []byte(reqStr)

	res, err := notionApi.Post("/pages", body, wrappers.ApiPostRequestOptions{})

	if err != nil {
		return item, err
	}

	json.NewDecoder(res.Body).Decode(&item)

	return item, nil
}
