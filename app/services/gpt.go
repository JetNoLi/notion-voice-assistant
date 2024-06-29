package services

import (
	"encoding/json"
	"fmt"

	"github.com/jetnoli/notion-voice-assistant/config/client"
	"github.com/jetnoli/notion-voice-assistant/wrappers/fetch"
)

type AssistRequest struct {
	Model     string          `json:"model"`
	MaxTokens int             `json:"max_tokens"`
	Messages  []AssistMessage `json:"messages"`
}

type AssistMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type AssistResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	SystemFingerprint interface{} `json:"system_fingerprint"`
}

func Assist(prompt string) (AssistResponse, error) {

	req := &AssistRequest{
		Model:     "gpt-3.5-turbo-16k",
		MaxTokens: 750,
		Messages:  []AssistMessage{{Role: "user", Content: prompt}},
	}

	body, err := json.Marshal(req)

	if err != nil {
		fmt.Println(err)
		return AssistResponse{}, err
	}

	res, err := client.OpenAiApi.Post("/chat/completions", body, fetch.ApiPostRequestOptions{})

	if err != nil {
		return AssistResponse{}, err
	}

	defer res.Body.Close()

	data := AssistResponse{}

	err = json.NewDecoder(res.Body).Decode(&data)

	if err != nil {
		return AssistResponse{}, err
	}

	return data, nil
}
