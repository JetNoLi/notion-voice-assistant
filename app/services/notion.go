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

type NotionTaskDB = notion.NotionTaskDB

// - Use GPT with prompt, to figure out which relation makes the most sense for each property
// - Use GPT with prompt, to figure out which value makes the most sense for each select
// - Use GPT with prompt, to generate a title and description
func GetAllRelatedPages(databaseId string) (any, error) {
	db, err := GetDatabaseById[NotionTaskDB](databaseId)

	if err != nil {
		return nil, err
	}

	relatedDBProps := make(map[string]RelatedDBPages)
	dbOptions := make(map[string]CreateDBOption)

	relatedDBs := []notion.NotionDBRelationProp{db.Properties.Categories, db.Properties.SubCategory, db.Properties.Project}

	for _, relatedDB := range relatedDBs {

		foreignDbPages, err := GetDatabasePagesById[notion.NotionPage[notion.NotionPageWithName]](relatedDB.Relation.DatabaseID)

		if err != nil {
			return nil, err
		}

		pageData := make([]PageNameAndID, 0)

		for _, page := range foreignDbPages.Results {
			pageData = append(pageData, PageNameAndID{
				Name: page.Properties.Name.Title[0].Text.Content,
				ID:   page.ID,
			})
		}

		relatedDBProps[relatedDB.Name] = RelatedDBPages{
			PageData: pageData,
			ID:       db.ID,
		}
	}

	for i, option := range append(append(db.Properties.Priority.Select.Options, db.Properties.Status.Status.Options...), db.Properties.Tags.MultiSelect.Options...) {

		if i < len(db.Properties.Priority.Select.Options) {
			dbOptions["Priority"] = CreateDBOption{
				Type:    db.Properties.Priority.Type,
				Options: append(dbOptions["Priority"].Options, option.Name),
			}
		} else if i < len(db.Properties.Status.Status.Options) {
			dbOptions["Status"] = CreateDBOption{
				Type:    db.Properties.Status.Type,
				Options: append(dbOptions["Status"].Options, option.Name),
			}
		} else if i < len(db.Properties.Tags.MultiSelect.Options) {
			dbOptions["Tags"] = CreateDBOption{
				Type:    db.Properties.Tags.Type,
				Options: append(dbOptions["Tags"].Options, option.Name),
			}
		}
	}

	req := notion.NotionCreateTaskRequestBuilder{}

	req.Add("db", db.ID)
	req.Add("name", "Page with Builder")
	req.Add("categories", relatedDBProps["Categories"].PageData[0].ID)
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
	}{Response: resp, Options: dbOptions, ForeignProps: relatedDBProps, Request: string(jsonReq)}, nil
}
