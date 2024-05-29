package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/jetnoli/notion-voice-assistant/config/client"
)

//TODO:
// - Respond Correctly
// - Work With the Service
// - Cleanup Code

type TranscriptionResponse struct {
	Message string `json:"message"`
	Result  string `json:"result"`
}

func Transcribe(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	r.ParseMultipartForm(10 << 20) // Max size == 10 MB

	file, handler, err := r.FormFile("audio_file")

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf(`{
			"title":"cannot parse audio file",
			"message": "%s"
		}`, err.Error())))
		return
	}

	defer file.Close()

	// Create a buffer to write the multipart form data
	var multiPartFileBuffer bytes.Buffer
	writer := multipart.NewWriter(&multiPartFileBuffer)

	// Create a form file field and copy the file content to it
	part, err := writer.CreateFormFile("file", handler.Filename)

	if err != nil {
		http.Error(w, "Error creating form file", http.StatusInternalServerError)
		return
	}
	_, err = io.Copy(part, file)
	if err != nil {
		http.Error(w, "Error copying file content", http.StatusInternalServerError)
		return
	}

	// Close the multipart writer to set the ending boundary
	err = writer.Close()
	if err != nil {
		http.Error(w, "Error closing writer", http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest("POST", client.WhisperApi.BaseUrl+"/transcribe/", &multiPartFileBuffer)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf(`{
			"title":"cannot prepare for transcription",
			"message": "%s"
		}`, err.Error())))
		return
	}

	// Set the Content-Type header to multipart/form-data with the boundary
	req.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := client.WhisperApi.Client.Do(req)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf(`{
			"title":"cannot transcribe audio file",
			"message": "%s"
		}`, err.Error())))
		return
	}

	defer res.Body.Close()

	if err != nil {

		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf(`{
			"title":"cannot str body",
			"message": "%s"
			}`, err.Error())))
		return
	}

	var data TranscriptionResponse

	err = json.NewDecoder(res.Body).Decode(&data)

	if err != nil {

		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf(`{
			"title":"cannot decode transcription",
			"data":"%s",
			"message": "%s"
			}`, string(data.Message), err.Error())))
		return
	}

	err = json.NewEncoder(w).Encode(data)

	if err != nil {

		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf(`{
			"title":"cannot encode transcription",
			"message": "%s"
			}`, err.Error())))
		return
	}
}
