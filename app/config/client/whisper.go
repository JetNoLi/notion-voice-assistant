package client

import (
	"net/http"

	"github.com/jetnoli/notion-voice-assistant/config"
	"github.com/jetnoli/notion-voice-assistant/wrappers/fetch"
)

var WhisperApi = fetch.Api{
	BaseUrl: config.WhisperApiUrl,
	Client:  &http.Client{},
	Headers: map[string]string{},
}
