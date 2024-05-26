package client

import (
	"net/http"

	"github.com/jetnoli/notion-voice-assistant/config"
	"github.com/jetnoli/notion-voice-assistant/wrappers"
)

var WhisperApi = wrappers.Api{
	BaseUrl: config.WhisperApiUrl,
	Client:  &http.Client{},
	Headers: map[string]string{},
}
