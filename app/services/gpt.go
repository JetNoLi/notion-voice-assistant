package services

import (
	"fmt"
	"io"

	"github.com/jetnoli/notion-voice-assistant/config/client"
	"github.com/jetnoli/notion-voice-assistant/wrappers"
)

func Assist(prompt string) (string, error) {

	fmt.Printf("This is the prompt:  %s", prompt)

	body := []byte(`{
		"model": "gpt-3.5-turbo-16k",
		"max_tokens": 1,
		"messages": 
			{
				"role": "user",
				"content": "%s"
			}
		
	}`)

	fmt.Println(string(body))

	res, err := client.OpenAiApi.Post("/chat/completions", body, wrappers.ApiPostRequestOptions{})

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
