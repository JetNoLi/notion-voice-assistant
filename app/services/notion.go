package services

import (
	"encoding/json"
	"fmt"

	"github.com/jetnoli/notion-voice-assistant/config/client"
	"github.com/jetnoli/notion-voice-assistant/wrappers/fetch"
)

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

	res, err := client.NotionApi.Post("/search", body, fetch.ApiPostRequestOptions{})

	if err != nil {
		return data, err
	}

	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&data)
	return data, err
}

func GetDatabaseById[T any](id string) (data T, err error) {
	res, err := client.NotionApi.Get("/databases/"+id, fetch.ApiGetRequestOptions{})

	if err != nil {
		return data, err
	}

	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&data)

	return data, err
}

type ItemData struct {
	Title       string
	StartDate   string //TODO: Look into Go Dates
	Status      string
	Tags        []string
	Project     string
	Category    string
	SubCategory string
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

	res, err := client.NotionApi.Post("/pages", body, fetch.ApiPostRequestOptions{})

	if err != nil {
		return item, err
	}

	json.NewDecoder(res.Body).Decode(&item)

	return item, nil
}

func Test(db NotionDB) {
	// Fetch Task DB
	// Map DB to Object
	// Fetch All Relations

	// Get Audio
	// Pass Info To Chat GPT
	// Determine Best Fit for Each Field

	// Create Item in DB
	// Create a Type for Accepting Data
	// Marshal Data to JSON
	// Create Task

	// Validate
}
