package client

import (
	"net/http"

	"github.com/jetnoli/notion-voice-assistant/config"
	"github.com/jetnoli/notion-voice-assistant/wrappers/fetch"
)

var OpenAiApi = fetch.Api{
	BaseUrl: config.OpenAiApiUrl + "/v1",
	Client:  &http.Client{},
	Headers: map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + config.OpenAiApiKey,
	},
}
