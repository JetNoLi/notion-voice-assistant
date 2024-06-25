package services

import (
	"encoding/json"
	"fmt"
	"time"
)

type NotionTask struct {
	Archived  bool        `json:"archived"`
	Cover     interface{} `json:"cover"`
	CreatedBy struct {
		ID     string `json:"id"`
		Object string `json:"object"`
	} `json:"created_by"`
	CreatedTime time.Time     `json:"created_time"`
	Description []interface{} `json:"description"`
	Icon        struct {
		Emoji string `json:"emoji"`
		Type  string `json:"type"`
	} `json:"icon"`
	ID           string `json:"id"`
	InTrash      bool   `json:"in_trash"`
	IsInline     bool   `json:"is_inline"`
	LastEditedBy struct {
		ID     string `json:"id"`
		Object string `json:"object"`
	} `json:"last_edited_by"`
	LastEditedTime time.Time `json:"last_edited_time"`
	Object         string    `json:"object"`
	Parent         struct {
		PageID string `json:"page_id"`
		Type   string `json:"type"`
	} `json:"parent"`
	Properties struct {
		Categories struct {
			ID       string `json:"id"`
			Name     string `json:"name"`
			Relation struct {
				DatabaseID   string `json:"database_id"`
				DualProperty struct {
					SyncedPropertyID   string `json:"synced_property_id"`
					SyncedPropertyName string `json:"synced_property_name"`
				} `json:"dual_property"`
				Type string `json:"type"`
			} `json:"relation"`
			Type string `json:"type"`
		} `json:"Categories"`
		CompleteTask struct {
			Button struct {
			} `json:"button"`
			ID   string `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"Complete Task"`
		CreatedTime struct {
			CreatedTime struct {
			} `json:"created_time"`
			ID   string `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"Created time"`
		DurationInMinutes struct {
			Formula struct {
				Expression string `json:"expression"`
			} `json:"formula"`
			ID   string `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"Duration in Minutes"`
		Goals struct {
			ID       string `json:"id"`
			Name     string `json:"name"`
			Relation struct {
				DatabaseID   string `json:"database_id"`
				DualProperty struct {
					SyncedPropertyID   string `json:"synced_property_id"`
					SyncedPropertyName string `json:"synced_property_name"`
				} `json:"dual_property"`
				Type string `json:"type"`
			} `json:"relation"`
			Type string `json:"type"`
		} `json:"Goals"`
		LastEditedTime struct {
			ID             string `json:"id"`
			LastEditedTime struct {
			} `json:"last_edited_time"`
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"Last edited time"`
		Name struct {
			ID    string `json:"id"`
			Name  string `json:"name"`
			Title struct {
			} `json:"title"`
			Type string `json:"type"`
		} `json:"Name"`
		Note struct {
			ID       string `json:"id"`
			Name     string `json:"name"`
			Relation struct {
				DatabaseID   string `json:"database_id"`
				DualProperty struct {
					SyncedPropertyID   string `json:"synced_property_id"`
					SyncedPropertyName string `json:"synced_property_name"`
				} `json:"dual_property"`
				Type string `json:"type"`
			} `json:"relation"`
			Type string `json:"type"`
		} `json:"Note"`
		Priority struct {
			ID     string `json:"id"`
			Name   string `json:"name"`
			Select struct {
				Options []struct {
					Color       string      `json:"color"`
					Description interface{} `json:"description"`
					ID          string      `json:"id"`
					Name        string      `json:"name"`
				} `json:"options"`
			} `json:"select"`
			Type string `json:"type"`
		} `json:"Priority"`
		Project struct {
			ID       string `json:"id"`
			Name     string `json:"name"`
			Relation struct {
				DatabaseID   string `json:"database_id"`
				DualProperty struct {
					SyncedPropertyID   string `json:"synced_property_id"`
					SyncedPropertyName string `json:"synced_property_name"`
				} `json:"dual_property"`
				Type string `json:"type"`
			} `json:"relation"`
			Type string `json:"type"`
		} `json:"Project"`
		Resource struct {
			ID       string `json:"id"`
			Name     string `json:"name"`
			RichText struct {
			} `json:"rich_text"`
			Type string `json:"type"`
		} `json:"Resource"`
		StartDate struct {
			Date struct {
			} `json:"date"`
			ID   string `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"Start Date"`
		Status struct {
			ID     string `json:"id"`
			Name   string `json:"name"`
			Status struct {
				Groups []struct {
					Color     string   `json:"color"`
					ID        string   `json:"id"`
					Name      string   `json:"name"`
					OptionIds []string `json:"option_ids"`
				} `json:"groups"`
				Options []struct {
					Color       string      `json:"color"`
					Description interface{} `json:"description"`
					ID          string      `json:"id"`
					Name        string      `json:"name"`
				} `json:"options"`
			} `json:"status"`
			Type string `json:"type"`
		} `json:"Status"`
		SubCategory struct {
			ID       string `json:"id"`
			Name     string `json:"name"`
			Relation struct {
				DatabaseID     string `json:"database_id"`
				SingleProperty struct {
				} `json:"single_property"`
				Type string `json:"type"`
			} `json:"relation"`
			Type string `json:"type"`
		} `json:"Sub Category"`
		Tags struct {
			ID          string `json:"id"`
			MultiSelect struct {
				Options []struct {
					Color       string      `json:"color"`
					Description interface{} `json:"description"`
					ID          string      `json:"id"`
					Name        string      `json:"name"`
				} `json:"options"`
			} `json:"multi_select"`
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"Tags"`
	} `json:"properties"`
	PublicURL interface{} `json:"public_url"`
	RequestID string      `json:"request_id"`
	Title     []struct {
		Annotations struct {
			Bold          bool   `json:"bold"`
			Code          bool   `json:"code"`
			Color         string `json:"color"`
			Italic        bool   `json:"italic"`
			Strikethrough bool   `json:"strikethrough"`
			Underline     bool   `json:"underline"`
		} `json:"annotations"`
		Href      interface{} `json:"href"`
		PlainText string      `json:"plain_text"`
		Text      struct {
			Content string      `json:"content"`
			Link    interface{} `json:"link"`
		} `json:"text"`
		Type string `json:"type"`
	} `json:"title"`
	URL string `json:"url"`
}

type NotionDB struct {
	CreatedBy struct {
		ID     string `json:"id"`
		Object string `json:"object"`
	} `json:"created_by"`
	LastEditedBy struct {
		ID     string `json:"id"`
		Object string `json:"object"`
	} `json:"last_edited_by"`
	CreatedTime time.Time `json:"created_time"`
	Icon        struct {
		Emoji string `json:"emoji"`
		Type  string `json:"type"`
	} `json:"icon"`

	Archived       bool                        `json:"archived"`
	Cover          interface{}                 `json:"cover"`
	Description    []interface{}               `json:"description"`
	ID             string                      `json:"id"`
	InTrash        bool                        `json:"in_trash"`
	IsInline       bool                        `json:"is_inline"`
	LastEditedTime time.Time                   `json:"last_edited_time"`
	Object         string                      `json:"object"`
	Parent         NotionDBParent              `json:"parent"`
	Properties     map[string]NotionDBProperty `json:"properties"`
	PublicURL      interface{}                 `json:"public_url"`
	RequestID      string                      `json:"request_id"`
	Title          []NotionTitleValue          `json:"title"`
	URL            string                      `json:"url"`
}

type NotionTitleValue struct {
	Annotations struct {
		Bold          bool   `json:"bold"`
		Code          bool   `json:"code"`
		Color         string `json:"color"`
		Italic        bool   `json:"italic"`
		Strikethrough bool   `json:"strikethrough"`
		Underline     bool   `json:"underline"`
	} `json:"annotations"`
	Href      interface{} `json:"href"`
	PlainText string      `json:"plain_text"`
	Text      struct {
		Content string      `json:"content"`
		Link    interface{} `json:"link"`
	} `json:"text"`
	Type string `json:"type"`
}

type NotionDBParent struct {
	PageID string `json:"page_id"`
	Type   string `json:"type"`
}

type NotionDBProperty struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value any    `json:"data"`
}

// UnmarshalJSON custom unmarshaling for NotionProperty to handle dynamic structures.
func (np *NotionDBProperty) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	np.ID = raw["id"].(string)
	np.Name = raw["name"].(string)
	np.Type = raw["type"].(string)

	var val any = nil

	// Store the raw value in the Value field for further processing
	if np.Type == "multi_select" {
		val = &NotionDBMultiSelectProp{}
	} else if np.Type == "status" {
		val = &NotionDBStatusProp{}
	} else if np.Type == "select" {
		val = &NotionDBSelectProp{}
	} else if np.Type == "relation" {
		val = &NotionDBRelationProp{}
	} else if np.Type == "date" {
		val = &NotionDBDateProp{}
	} else if np.Type == "title" {
		val = &NotionDBTitleProp{}
	}

	if val == nil {
		fmt.Println("skipping", np.Type)
		np.Value = raw[np.Type]
		return nil
	}

	if err := json.Unmarshal(data, val); err != nil {
		fmt.Println("failed type assertion status", err.Error())
		np.Value = raw[np.Type]
		return err
	}

	np.Value = val
	return nil
}

type NotionDBRelationProp struct {
	Relation NotionDBRelationPropValue `json:"relation"`
}

type NotionDBRelationPropValue struct {
	DatabaseID   string `json:"database_id"`
	DualProperty struct {
		SyncedPropertyID   string `json:"synced_property_id"`
		SyncedPropertyName string `json:"synced_property_name"`
	} `json:"dual_property"`
	Type string `json:"type"`
}

type NotionDBDateProp struct {
	Date NotionDatePropValue
}

type NotionDatePropValue struct {
	Start    string `json:"start"`
	End      string `json:"end,omitempty"`
	TimeZone string `json:"time_zone,omitempty"` //TODO: Make sure adding date works?
}

type NotionDBMultiSelectProp struct {
	MultiSelect NotionDBSelectValue `json:"multi_select"`
}

type NotionDBSelectProp struct {
	Select NotionDBSelectValue `json:"select"`
}

type NotionDBSelectValue struct {
	Options []NotionDBSelectPropOptions `json:"options"`
}

type NotionDBTitleProp struct {
}

type NotionDBRichTextProp struct{}

type NotionDBStatusProp struct {
	Status struct {
		Groups []struct {
			Color     string   `json:"color"`
			ID        string   `json:"id"`
			Name      string   `json:"name"`
			OptionIds []string `json:"option_ids"`
		} `json:"groups"`
		Options []NotionDBSelectPropOptions `json:"options"`
	} `json:"status"`
}

type NotionDBSelectPropOptions struct {
	Color       string      `json:"color"`
	Description interface{} `json:"description"`
	ID          string      `json:"id"`
	Name        string      `json:"name"`
}

// * Categories - Relation
// * Sub Category - Relation
// * Goals - Relation
// * Note - Relation
// * Project - Relation
// * Name - Title
// * Priority - Select
// * Start Date - Date
// * Status - Status
// * Tags - Multi Select

// Skip On DB Property Types
// * Resource - Rich Text v2 TODO:
// * formula
// * rollup
// * last_edited_time
// * created_time
// * button

type NotionPage[Props any] struct {
	Object         string    `json:"object"`
	ID             string    `json:"id"`
	CreatedTime    time.Time `json:"created_time"`
	LastEditedTime time.Time `json:"last_edited_time"`
	CreatedBy      struct {
		Object string `json:"object"`
		ID     string `json:"id"`
	} `json:"created_by"`
	LastEditedBy struct {
		Object string `json:"object"`
		ID     string `json:"id"`
	} `json:"last_edited_by"`
	Cover interface{} `json:"cover"`
	Icon  struct {
		Type  string `json:"type"`
		Emoji string `json:"emoji"`
	} `json:"icon"`
	Parent struct {
		Type       string `json:"type"`
		DatabaseID string `json:"database_id"`
	} `json:"parent"`
	Archived   bool        `json:"archived"`
	InTrash    bool        `json:"in_trash"`
	Properties Props       `json:"properties"` //TODO: Add Props Typing
	URL        string      `json:"url"`
	PublicURL  interface{} `json:"public_url"`
}

type NotionPageRelationPropValue struct {
	ID string `json:"id"`
}

type NotionPageProp[T any] struct {
	ID             string      `json:"id"`
	Type           string      `json:"type"`
	HasMore        bool        `json:"has_more"`
	Value          T           `json:"data"`
	NextCursor     interface{} `json:"next_cursor"`
	PageOrDatabase struct {
	} `json:"page_or_database"`
	DeveloperSurvey string `json:"developer_survey"`
	RequestID       string `json:"request_id"`
}

type NotionPageRelationProp = NotionPageProp[[]NotionPageRelationPropValue]
type NotionPageTitleProp = NotionPageProp[[]NotionTitleValue]

type NotionCreateTaskRequest = NotionCreatePageRequest[NotionTaskDBPageProps]

type NotionCreatePageRequest[Props any] struct {
	Parent struct {
		DatabaseID string `json:"database_id"`
	} `json:"parent"`
	Icon *struct {
		Emoji string `json:"emoji"`
	} `json:"icon,omitempty"`
	Cover *struct {
		External struct {
			URL string `json:"url"`
		} `json:"external"`
	} `json:"cover,omitempty"`
	Properties Props `json:"properties"`
}

func (npr *NotionCreatePageRequest[NotionTaskDBPageProps]) ToJSON() ([]byte, error) {
	reqStr, err := json.Marshal(npr)

	return reqStr, err
}

type NotionPageCreateNameProp struct {
	Title []NotionText `json:"title"`
}

type NotionText struct {
	Text NotionContent `json:"text"`
}

type NotionContent struct {
	Content string `json:"content"`
}

type NotionPageCreateRichTextProp struct {
	RichText []NotionText `json:"rich_text"`
}

type NotionPageCreateMultiSelectProp struct {
	MultiSelect []NotionMultiSelect `json:"multi_select"`
}

type NotionMultiSelect struct {
	Name *string `json:"name"`
}

type NotionPageCreateRelationProp struct {
	Relation []NotionRelation `json:"relation"`
}

type NotionRelation struct {
	ID string `json:"id"`
}

type NotionPageCreateSelectProp struct {
	Select NotionSelect `json:"select"`
}

type NotionSelect struct {
	Name string `json:"name"`
}

type NotionPageCreateStatusProp struct {
	Status NotionStatus `json:"status"`
}

type NotionStatus struct {
	Name string `json:"name"`
}

type NotionPageCreateDateProp struct {
	Date NotionDatePropValue `json:"date"`
}

type NotionTaskDBPageProps struct {
	Name        *NotionPageCreateNameProp        `json:"Name"`
	Resource    *NotionPageCreateRichTextProp    `json:"Resource,omitempty"`
	Tags        *NotionPageCreateMultiSelectProp `json:"Tags,omitempty"`
	Categories  *NotionPageCreateRelationProp    `json:"Categories,omitempty"`
	SubCategory *NotionPageCreateRelationProp    `json:"Sub Category,omitempty"`
	Goals       *NotionPageCreateRelationProp    `json:"Goals,omitempty"`
	Note        *NotionPageCreateRelationProp    `json:"Note,omitempty"`
	Project     *NotionPageCreateRelationProp    `json:"Project,omitempty"`
	Priority    *NotionPageCreateSelectProp      `json:"Priority,omitempty"`
	StartDate   *NotionPageCreateDateProp        `json:"Start Date,omitempty"`
	Status      *NotionPageCreateStatusProp      `json:"Status,omitempty"`
}
