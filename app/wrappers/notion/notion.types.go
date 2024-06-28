package notion

import (
	"encoding/json"
	"time"
)

//TODO:
// Create Concrete Types for DB
// Break Function Down into Pieces
// Figure out GPT Prompt

type NotionDB[Props any] struct {
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

	Archived       bool               `json:"archived"`
	Cover          interface{}        `json:"cover"`
	Description    []interface{}      `json:"description"`
	ID             string             `json:"id"`
	InTrash        bool               `json:"in_trash"`
	IsInline       bool               `json:"is_inline"`
	LastEditedTime time.Time          `json:"last_edited_time"`
	Object         string             `json:"object"`
	Parent         NotionDBParent     `json:"parent"`
	Properties     Props              `json:"properties"`
	PublicURL      interface{}        `json:"public_url"`
	RequestID      string             `json:"request_id"`
	Title          []NotionTitleValue `json:"title"`
	URL            string             `json:"url"`
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

type NotionTaskDBProps struct {
	Categories  NotionDBRelationProp    `json:"Categories"`
	Name        NotionDBTitleProp       `json:"Name"`
	Note        NotionDBRelationProp    `json:"Note"`
	Priority    NotionDBSelectProp      `json:"Priority"`
	Project     NotionDBRelationProp    `json:"Project"`
	Resource    NotionDBRichTextProp    `json:"Resource"`
	Status      NotionDBStatusProp      `json:"Status"`
	SubCategory NotionDBRelationProp    `json:"Sub Category"`
	Tags        NotionDBMultiSelectProp `json:"Tags"`
}

type NotionPageWithName struct {
	Name NotionPageCreateNameProp `json:"Name"`
}

type NotionTaskDB = NotionDB[NotionTaskDBProps]

type NotionDBRelationProp struct {
	ID       string                    `json:"id"`
	Name     string                    `json:"name"`
	Type     string                    `json:"type"`
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
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type NotionDatePropValue struct {
	Start    string `json:"start"`
	End      string `json:"end,omitempty"`
	TimeZone string `json:"time_zone,omitempty"` //TODO: Make sure adding date works?
}

type NotionDBMultiSelectProp struct {
	ID          string              `json:"id"`
	Name        string              `json:"name"`
	Type        string              `json:"type"`
	MultiSelect NotionDBSelectValue `json:"multi_select"`
}

type NotionDBSelectProp struct {
	ID     string              `json:"id"`
	Name   string              `json:"name"`
	Type   string              `json:"type"`
	Select NotionDBSelectValue `json:"select"`
}

type NotionDBSelectValue struct {
	Options []NotionDBSelectPropOptions `json:"options"`
}

type NotionDBTitleProp struct {
}

type NotionDBRichTextProp struct{}

type NotionDBStatusProp struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
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
	Properties Props       `json:"properties"`
	URL        string      `json:"url"`
	PublicURL  interface{} `json:"public_url"`
}

type NotionPageRelationPropValue struct {
	ID string `json:"id"`
}

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

func (npr *NotionCreatePageRequest[any]) ToJSON() ([]byte, error) {
	reqStr, err := json.Marshal(npr)

	return reqStr, err
}

type NotionRequestInterface interface {
	ToJSON() ([]byte, error)
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
