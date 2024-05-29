package services

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/jetnoli/notion-voice-assistant/config/client"
)

type TranscriptionResponse struct {
	Message string `json:"message"`
	Result  string `json:"result"`
}

func TranscribeFile(multiPartFileBuffer *bytes.Buffer, contentType string) (data TranscriptionResponse, err error) {
	req, err := http.NewRequest("POST", client.WhisperApi.BaseUrl+"/transcribe/", multiPartFileBuffer)

	if err != nil {
		return data, err
	}
	defer req.Body.Close()

	// Set the Content-Type header to multipart/form-data with the boundary
	req.Header.Set("Content-Type", contentType)

	res, err := client.WhisperApi.Client.Do(req)

	if err != nil {
		return data, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&data)

	return data, err
}
