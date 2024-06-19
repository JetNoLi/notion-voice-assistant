package services

import "time"

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

	Archived       bool           `json:"archived"`
	Cover          interface{}    `json:"cover"`
	Description    []interface{}  `json:"description"`
	ID             string         `json:"id"`
	InTrash        bool           `json:"in_trash"`
	IsInline       bool           `json:"is_inline"`
	LastEditedTime time.Time      `json:"last_edited_time"`
	Object         string         `json:"object"`
	Parent         NotionDBParent `json:"parent"`
	Properties     map[string]NotionDBProperty
	PublicURL      interface{}   `json:"public_url"`
	RequestID      string        `json:"request_id"`
	Title          NotionDBTitle `json:"title"`
	URL            string        `json:"url"`
}

type NotionDBTitle struct {
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
	ID    string      `json:"id"`
	Name  string      `json:"name"`
	Type  string      `json:"type"`
	Value interface{} `json:"-"`
}

type NotionDBRelationProp struct {
	DatabaseID   string `json:"database_id"`
	DualProperty struct {
		SyncedPropertyID   string `json:"synced_property_id"`
		SyncedPropertyName string `json:"synced_property_name"`
	} `json:"dual_property"`
	Type string `json:"type"`
}

type NotionDBMultiSelectProp struct {
	Options []NotionDBSelectPropOptions `json:"options"`
}

type NotionDBSelectProp struct {
	Options []NotionDBSelectPropOptions `json:"options"`
}

type NotionDBTitleProp struct{}
type NotionDBDateProp struct{}
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

// Pages

// "properties": {
// 	"Name": {
// 		"title": [
// 			{
// 				"text": {
// 					"content": "%s"
// 				}
// 			}
// 		]
// 	},
// 	"Tags": {
// 		"multi_select": [
// 			{
// 				"name": "Workout"
// 			}
// 		]
// 	}

// }
// * Categories - Relation
// * Sub Category - Relation
// * Goals - Relation
// * Note - Relation
// * Project - Relation
// * Name - Title
// * Priority - Select
// * Resource - Rich Text
// * Start Date - Date
// * Status - Status
// * Tags - Multi Select

// Skip On DB Property Types
// * formula
// * rollup
// * last_edited_time
// * created_time
// * button

type CreateNotionDBPageRequest struct {
	Parent struct {
		DatabaseID string `json:"database_id"`
	} `json:"parent"`
	Icon struct {
		Emoji string `json:"emoji"`
	} `json:"icon"`
	Properties map[string]struct {
		Key map[string]struct {
		} `json:"-"`
	}
}
