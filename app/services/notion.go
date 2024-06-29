package services

import (
	"encoding/json"
	"fmt"
	"strings"

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

func CreateDatabaseItem[Res any, Req notion.RequestInterface](databaseId string, itemData *Req) (item *Res, err error) {
	body, err := (*itemData).ToJSON()

	fmt.Printf("my data \n %#v\n", *itemData)

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

type NotionDBOption struct {
	Type    string
	Options []string
}

type NotionTaskDB = notion.TaskDB
type NotionTaskPage = notion.Page[notion.TaskDBPageProps]

type NotionRelatedPageContent struct {
	Options      map[string]NotionDBOption
	ForeignProps map[string]RelatedDBPages
}

func GetAllRelatedPages(databaseId string) (*NotionRelatedPageContent, error) {
	db, err := GetDatabaseById[NotionTaskDB](databaseId)

	if err != nil {
		return nil, err
	}

	relatedDBProps := make(map[string]RelatedDBPages)
	dbOptions := make(map[string]NotionDBOption)

	relatedDBs := []notion.DBRelationProp{db.Properties.Categories, db.Properties.SubCategory, db.Properties.Project}

	for _, relatedDB := range relatedDBs {

		foreignDbPages, err := GetDatabasePagesById[notion.Page[notion.PageWithName]](relatedDB.Relation.DatabaseID)

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
			dbOptions["Priority"] = NotionDBOption{
				Type:    db.Properties.Priority.Type,
				Options: append(dbOptions["Priority"].Options, option.Name),
			}
		} else if i < len(db.Properties.Status.Status.Options) {
			dbOptions["Status"] = NotionDBOption{
				Type:    db.Properties.Status.Type,
				Options: append(dbOptions["Status"].Options, option.Name),
			}
		} else if i < len(db.Properties.Tags.MultiSelect.Options) {
			dbOptions["Tags"] = NotionDBOption{
				Type:    db.Properties.Tags.Type,
				Options: append(dbOptions["Tags"].Options, option.Name),
			}
		}
	}

	return &NotionRelatedPageContent{Options: dbOptions, ForeignProps: relatedDBProps}, nil
}

// - Use GPT with prompt, to figure out which relation makes the most sense for each property
// - Use GPT with prompt, to figure out which value makes the most sense for each select
// - Use GPT with prompt, to generate a title and description
type TaskRelation struct {
	Name   string `json:"Name"`
	PageID string `json:"PageID"`
	DBID   string `json:"DBID"`
}

type TaskData struct {
	Name      string                    `json:"name"`
	Options   map[string]any            `json:"options"`   // string | []string
	Relations map[string][]TaskRelation `json:"relations"` // TaskRelation | []TaskRelation
}

func CreatePageFromRelatedContent(databaseId string, relatedContent *NotionRelatedPageContent, prompt string) (any, error) {
	data, err := Assist(notion.CreatePrompt(notion.CreatePromptArgs{
		Prompt:       prompt,
		DBFields:     notion.TaskDBPageProps{},
		DBRelations:  relatedContent.ForeignProps,
		Options:      relatedContent.Options,
		ReturnStruct: TaskData{},
	}))

	if err != nil {
		return nil, err
	}

	fmt.Println("starting marshal")
	taskData := &TaskData{}

	if err = json.Unmarshal([]byte(data.Choices[0].Message.Content), taskData); err != nil {
		return nil, fmt.Errorf("error converting %s : %#v", err.Error(), data.Choices[0].Message.Content)
	}

	req := notion.CreateTaskRequestBuilder{}

	req.Add("db", databaseId)
	req.Add("name", taskData.Name)

	for key, value := range taskData.Options {
		fmt.Println("adding", key, value)
		val, ok := value.(string)

		if ok {
			req.Add(strings.ToLower(key), val)
			continue
		}

		vals, ok := value.([]any)

		if !ok {
			return nil, fmt.Errorf("invalid format for key "+key+"with val %#v", value)
		}

		for _, val := range vals {
			str, ok := val.(string)

			if ok {
				req.Add(strings.ToLower(key), str)
				continue
			}

			return nil, fmt.Errorf("error converting option to string in multiselect %v, %v", val, vals)
		}

	}

	for key, value := range taskData.Relations {
		fmt.Printf("adding %v %#v\n", key, value)

		for _, val := range value {
			req.Add(strings.ToLower(key), val.PageID)
		}

	}

	reqJson, err := req.Request()

	if err != nil {
		fmt.Println("error with reqest", err.Error())
	}

	resp, err := CreateDatabaseItem[any](databaseId, &reqJson)

	if err != nil {
		fmt.Println("error creating task", err.Error(), reqJson.Properties)
	}

	return resp, err

}
