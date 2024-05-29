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
	var b bytes.Buffer
	writer := multipart.NewWriter(&b)

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

	req, err := http.NewRequest("POST", client.WhisperApi.BaseUrl+"/transcribe/", &b)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf(`{
			"title":"cannot prepare for transcription",
			"message": "%s"
		}`, err.Error())))
		return
	}

	// Set the appropriate headers
	// req.Header.Set("Content-Type", handler.Header.Get("Content-Type"))
	// req.Header.Set("Content-Disposition", "form-data; name=\"file\"; filename=\""+handler.Filename+"\"")

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

	// splitFile := strings.Split(handler.Filename, ".")
	// fileType := splitFile[len(splitFile)-1]

	// tempFile, err := os.CreateTemp("audio", "nva-audio-*."+fileType)

	// if err != nil {
	// 	w.WriteHeader(400)
	// 	w.Write([]byte(fmt.Sprintf(`{
	// 		"title":"cannot parse audio file",
	// 		"message": "%s"
	// 	}`, err.Error())))
	// 	return
	// }

	// fileBytes, err := io.ReadAll(tempFile) //TODO: Try and pass just the file?

	// if err != nil {
	// 	w.WriteHeader(400)
	// 	w.Write([]byte(fmt.Sprintf(`{
	// 		"title":"cannot parse audio file",
	// 		"message": "%s"
	// 	}`, err.Error())))
	// 	return
	// }

	// // write this byte array to our temporary file
	// tempFile.Write(fileBytes)

	defer res.Body.Close()

	strBody, err := io.ReadAll(res.Body)

	if err != nil {

		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf(`{
			"title":"cannot str body",
			"message": "%s"
			}`, err.Error())))
		return
	}

	fmt.Println(string(strBody))

	var data any

	err = json.NewDecoder(res.Body).Decode(&data)

	if err != nil {

		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf(`{
			"title":"cannot decode transcription",
			"message": "%s"
			}`, err.Error())))
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

	w.Write([]byte(`{
		"message": "File Successfully Written"
	}`))

}
