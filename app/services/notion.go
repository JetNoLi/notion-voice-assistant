package services

import (
	"encoding/json"
	"fmt"

	"github.com/jetnoli/notion-voice-assistant/config/client"
	"github.com/jetnoli/notion-voice-assistant/wrappers/fetch"
	"github.com/jetnoli/notion-voice-assistant/wrappers/notion"
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

type NotionListResponse[T any] struct {
	Object  string `json:"object"`
	Results []T    `json:"results"`
}

func GetDatabasePagesById[T any](id string) (data NotionListResponse[T], err error) {
	res, err := client.NotionApi.Post("/databases/"+id+"/query", []byte(`{
		"sorts": [
        	{
       			"property": "Name",
        		"direction": "ascending"
        	}
    	]
	}`), fetch.ApiPostRequestOptions{})

	if err != nil {
		return data, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&data)

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

func CreateDatabaseItem[Res any, Req notion.NotionRequestInterface](databaseId string, itemData *Req) (item *Res, err error) {
	body, err := (*itemData).ToJSON()

	if err != nil {
		return item, err
	}

	res, err := client.NotionApi.Post("/pages", body, fetch.ApiPostRequestOptions{})

	if err != nil {
		return item, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&item)

	return item, err
}

func Test(db notion.NotionDB) {
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

// TODO: Make a struct { dbName, dbId, dbPageNames}
type PageNameAndID struct {
	Name string
	ID   string
}
type RelatedDBPages struct {
	ID       string
	PageData []PageNameAndID
}

type CreateDBOption struct {
	Type    string
	Options []string
}

// - Use GPT with prompt, to figure out which relation makes the most sense for each property
// - Use GPT with prompt, to figure out which value makes the most sense for each select
// - Use GPT with prompt, to generate a title and description
func GetAllRelatedPages(databaseId string) (any, error) {
	db, err := GetDatabaseById[notion.NotionDB](databaseId)

	if err != nil {
		return nil, err
	}

	foreignDBProps := make(map[string]RelatedDBPages)
	dbOptions := make(map[string]CreateDBOption)

	for _, prop := range db.Properties {
		//TODO: Make Switch Statement

		if prop.Type == "relation" {
			relation, ok := prop.Value.(*notion.NotionDBRelationProp)

			if !ok {
				fmt.Println("not ok", prop.Value)
				return nil, fmt.Errorf("error processing relation")
			}

			id := relation.Relation.DatabaseID //TODO: Add error handling

			fmt.Println("lloking up", id)

			foreignDbPages, err := GetDatabasePagesById[notion.NotionPage[map[string]any]](id)

			if err != nil {
				return nil, err
			}

			foreignDB, err := GetDatabaseById[notion.NotionDB](id)

			if err != nil {
				return nil, err
			}

			pageData := make([]PageNameAndID, 0)

			for _, page := range foreignDbPages.Results {
				for key, prop := range page.Properties {
					if key == "Name" {
						val, ok := prop.(map[string]any)

						if !ok {
							return nil, fmt.Errorf("invalid conversion on prop in page")
						}

						if val["title"] == nil {
							continue
						}

						titles, ok := val["title"].([]any)

						if !ok {
							return nil, fmt.Errorf("invalid conversion on prop in titles")
						}

						title, ok := titles[0].(map[string]any)

						if !ok {
							return nil, fmt.Errorf("invalid conversion on prop in title")
						}

						pageData = append(pageData, PageNameAndID{
							Name: title["plain_text"].(string),
							ID:   page.ID,
						})
					}
				}

			}

			foreignDBProps[foreignDB.Title[0].PlainText] = RelatedDBPages{
				PageData: pageData,
				ID:       id,
			}
		} else if prop.Type == "multi_select" {
			val, ok := prop.Value.(*notion.NotionDBMultiSelectProp)

			if !ok {
				fmt.Println("not ok", prop.Value)
				return nil, fmt.Errorf("error processing multi select")
			}

			options := make([]string, len(val.MultiSelect.Options))

			for i, option := range val.MultiSelect.Options {
				options[i] = option.Name
			}

			dbOptions[prop.Name] = CreateDBOption{
				Type:    prop.Type,
				Options: options,
			}
		} else if prop.Type == "status" {
			val, ok := prop.Value.(*notion.NotionDBStatusProp)

			if !ok {
				fmt.Println("not ok", prop.Value)
				return nil, fmt.Errorf("error processing status")
			}

			// TODO: Implement Map function
			options := make([]string, len(val.Status.Options))

			for i, option := range val.Status.Options {
				options[i] = option.Name
			}

			dbOptions[prop.Name] = CreateDBOption{
				Type:    prop.Type,
				Options: options,
			}

		} else if prop.Type == "select" {
			val, ok := prop.Value.(*notion.NotionDBSelectProp)

			if !ok {
				fmt.Println("not ok", prop.Value)
				return nil, fmt.Errorf("error processing select")
			}

			options := make([]string, len(val.Select.Options))

			for i, option := range val.Select.Options {
				options[i] = option.Name
			}

			dbOptions[prop.Name] = CreateDBOption{
				Type:    prop.Type,
				Options: options,
			}
		} else if prop.Type == "date" {
			_, ok := prop.Value.(*notion.NotionDBDateProp)

			if !ok {
				fmt.Println("not ok", prop.Value)
				return nil, fmt.Errorf("error processing date")
			}

		} else if prop.Type == "title" {
			_, ok := prop.Value.(*notion.NotionDBTitleProp)

			if !ok {
				fmt.Println("not ok", prop.Value)
				return nil, fmt.Errorf("error processing title")
			}

		} else {
			fmt.Println("skip", prop.Type)
		}
	}

	req := notion.NotionCreateTaskRequestBuilder{}

	req.Add("db", db.ID)
	req.Add("name", "Page with Builder")
	req.Add("categories", foreignDBProps["Categories"].PageData[0].ID)
	req.Add("status", dbOptions["Status"].Options[0])

	resp, err := CreateDatabaseItem[any](databaseId, &req.Builder.Req)

	if err != nil {
		return nil, err
	}

	jsonReq, err := req.Builder.Req.ToJSON()

	if err != nil {
		fmt.Println("could not parse to JSON")
	}

	return struct {
		Response     any
		Options      any
		ForeignProps any
		Request      any
	}{Response: resp, Options: dbOptions, ForeignProps: foreignDBProps, Request: string(jsonReq)}, nil
}
