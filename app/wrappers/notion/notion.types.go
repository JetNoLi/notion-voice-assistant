package notion

import (
	"encoding/json"
	"time"
)

//TODO:
// Create Concrete Types for DB
// Break Function Down into Pieces
// Figure out GPT Prompt

type DB[Props any] struct {
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

	Archived       bool          `json:"archived"`
	Cover          interface{}   `json:"cover"`
	Description    []interface{} `json:"description"`
	ID             string        `json:"id"`
	InTrash        bool          `json:"in_trash"`
	IsInline       bool          `json:"is_inline"`
	LastEditedTime time.Time     `json:"last_edited_time"`
	Object         string        `json:"object"`
	Parent         DBParent      `json:"parent"`
	Properties     Props         `json:"properties"`
	PublicURL      interface{}   `json:"public_url"`
	RequestID      string        `json:"request_id"`
	Title          []TitleValue  `json:"title"`
	URL            string        `json:"url"`
}

type TitleValue struct {
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

type DBParent struct {
	PageID string `json:"page_id"`
	Type   string `json:"type"`
}

type TaskDBProps struct {
	Categories  DBRelationProp    `json:"Categories"`
	Name        DBTitleProp       `json:"Name"`
	Note        DBRelationProp    `json:"Note"`
	Priority    DBSelectProp      `json:"Priority"`
	Project     DBRelationProp    `json:"Project"`
	Resource    DBRichTextProp    `json:"Resource"`
	Status      DBStatusProp      `json:"Status"`
	SubCategory DBRelationProp    `json:"Sub Category"`
	Tags        DBMultiSelectProp `json:"Tags"`
}

type PageWithName struct {
	Name PageCreateNameProp `json:"Name"`
}

type TaskDB = DB[TaskDBProps]

type DBRelationProp struct {
	ID       string              `json:"id"`
	Name     string              `json:"name"`
	Type     string              `json:"type"`
	Relation DBRelationPropValue `json:"relation"`
}

type DBRelationPropValue struct {
	DatabaseID   string `json:"database_id"`
	DualProperty struct {
		SyncedPropertyID   string `json:"synced_property_id"`
		SyncedPropertyName string `json:"synced_property_name"`
	} `json:"dual_property"`
	Type string `json:"type"`
}

type DBDateProp struct {
	Date DatePropValue
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type DatePropValue struct {
	Start    string `json:"start"`
	End      string `json:"end,omitempty"`
	TimeZone string `json:"time_zone,omitempty"` //TODO: Make sure adding date works?
}

type DBMultiSelectProp struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Type        string        `json:"type"`
	MultiSelect DBSelectValue `json:"multi_select"`
}

type DBSelectProp struct {
	ID     string        `json:"id"`
	Name   string        `json:"name"`
	Type   string        `json:"type"`
	Select DBSelectValue `json:"select"`
}

type DBSelectValue struct {
	Options []DBSelectPropOptions `json:"options"`
}

type DBTitleProp struct {
}

type DBRichTextProp struct{}

type DBStatusProp struct {
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
		Options []DBSelectPropOptions `json:"options"`
	} `json:"status"`
}

type DBSelectPropOptions struct {
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

type Page[Props any] struct {
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

type PageRelationPropValue struct {
	ID string `json:"id"`
}

type CreateTaskRequest = CreatePageRequest[TaskDBPageProps]

type CreatePageRequest[Props any] struct {
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

func (npr *CreatePageRequest[any]) ToJSON() ([]byte, error) {
	reqStr, err := json.Marshal(npr)

	return reqStr, err
}

type RequestInterface interface {
	ToJSON() ([]byte, error)
}

type PageCreateNameProp struct {
	Title []Text `json:"title"`
}

type Text struct {
	Text Content `json:"text"`
}

type Content struct {
	Content string `json:"content"`
}

type PageCreateRichTextProp struct {
	RichText []Text `json:"rich_text"`
}

type PageCreateMultiSelectProp struct {
	MultiSelect []MultiSelect `json:"multi_select"`
}

type MultiSelect struct {
	Name *string `json:"name"`
}

type PageCreateRelationProp struct {
	Relation []Relation `json:"relation"`
}

type Relation struct {
	ID string `json:"id"`
}

type PageCreateSelectProp struct {
	Select Select `json:"select"`
}

type Select struct {
	Name string `json:"name"`
}

type PageCreateStatusProp struct {
	Status Status `json:"status"`
}

type Status struct {
	Name string `json:"name"`
}

type PageCreateDateProp struct {
	Date DatePropValue `json:"date"`
}

type TaskDBPageProps struct {
	Name        *PageCreateNameProp        `json:"Name"`
	Resource    *PageCreateRichTextProp    `json:"Resource,omitempty"`
	Tags        *PageCreateMultiSelectProp `json:"Tags,omitempty"`
	Categories  *PageCreateRelationProp    `json:"Categories,omitempty"`
	SubCategory *PageCreateRelationProp    `json:"Sub Category,omitempty"`
	Goals       *PageCreateRelationProp    `json:"Goals,omitempty"`
	Note        *PageCreateRelationProp    `json:"Note,omitempty"`
	Project     *PageCreateRelationProp    `json:"Project,omitempty"`
	Priority    *PageCreateSelectProp      `json:"Priority,omitempty"`
	StartDate   *PageCreateDateProp        `json:"Start Date,omitempty"`
	Status      *PageCreateStatusProp      `json:"Status,omitempty"`
}
