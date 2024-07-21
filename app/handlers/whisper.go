package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jetnoli/notion-voice-assistant/config"
	"github.com/jetnoli/notion-voice-assistant/services"
)

func Transcribe(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if err := r.ParseMultipartForm(config.MAX_FILE_SIZE_IN_MEMORY); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf(`{
			"title":"error processing file",
			"message": "%s"
			}`, err.Error())))
		return
	}

	file, header, err := r.FormFile("audio_file")

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf(`{
			"title":"cannot parse audio file",
			"message": "%s"
		}`, err.Error())))
		return
	}

	defer file.Close()

	multiPartFileBuffer, contentType, err := services.CreateBufferFromFormFile(file, *header)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf(`{
			"title":"cannot decode transcription",
			"message": "%s"
			}`, err.Error())))
		return
	}

	data, err := services.TranscribeFile(&multiPartFileBuffer, contentType)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf(`{
			"title":"cannot decode transcription",
			"data":"%s",
			"message": "%s"
			}`, string(data.Message), err.Error())))
		return
	}

	w.Header().Add("Content-Type", "application/json")

	if err = json.NewEncoder(w).Encode(data); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf(`{
			"title":"cannot encode transcription",
			"message": "%s"
			}`, err.Error())))
	}
}

func TranscribeResponse(r *http.Request) (services.TranscriptionResponse, error) {

	if err := r.ParseMultipartForm(config.MAX_FILE_SIZE_IN_MEMORY); err != nil {
		return services.TranscriptionResponse{}, err
	}

	file, header, err := r.FormFile("audio_file")

	if err != nil {
		return services.TranscriptionResponse{}, err
	}

	defer file.Close()

	multiPartFileBuffer, contentType, err := services.CreateBufferFromFormFile(file, *header)

	if err != nil {
		return services.TranscriptionResponse{}, err
	}

	return services.TranscribeFile(&multiPartFileBuffer, contentType)

}
