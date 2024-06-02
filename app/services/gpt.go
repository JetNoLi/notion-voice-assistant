package services

import (
	"fmt"
	"io"

	"github.com/jetnoli/notion-voice-assistant/config/client"
	"github.com/jetnoli/notion-voice-assistant/wrappers/fetch"
)

func Assist(prompt string) (string, error) {

	body := []byte(fmt.Sprintf(`{
		"model": "gpt-3.5-turbo-16k",
		"max_tokens": 45,
		"messages": [
			{
				"role": "user",
				"content": "%s"
			}
		]
	}`, prompt))

	res, err := client.OpenAiApi.Post("/chat/completions", body, fetch.ApiPostRequestOptions{})

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	return string(data), nil
}
