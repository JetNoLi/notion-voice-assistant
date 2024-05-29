package client

import (
	"net/http"

	"github.com/jetnoli/notion-voice-assistant/config"
	"github.com/jetnoli/notion-voice-assistant/wrappers"
)

var OpenAiApi = wrappers.Api{
	BaseUrl: config.OpenAiApiUrl + "/v1",
	Client:  &http.Client{},
	Headers: map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + config.OpenAiApiKey,
	},
}
