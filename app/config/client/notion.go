package client

import (
	"net/http"

	"github.com/jetnoli/notion-voice-assistant/config"
	"github.com/jetnoli/notion-voice-assistant/wrappers"
)

var NotionApi = wrappers.Api{
	BaseUrl: "https://api.notion.com/v1",
	Client:  &http.Client{},
	Headers: map[string]string{
		"Content-Type":   "application/json",
		"Authorization":  "Bearer " + config.NotionApiKey,
		"Notion-Version": "2022-06-28",
	},
}
