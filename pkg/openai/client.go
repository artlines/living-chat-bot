package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	httpClient = &http.Client{Timeout: 10 * time.Second}
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type ChatResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

func SendMessage(userInput string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Println("OPENAI_API_KEY not set in environment")
		return "", fmt.Errorf("OPENAI_API_KEY not set in environment")
	}

	openAIURL := os.Getenv("OPENAI_URL")
	if openAIURL == "" {
		log.Println("OPENAI_URL not set in environment")
		return "", fmt.Errorf("OPENAI_URL not set in environment")
	}

	httpClient = &http.Client{
		Timeout: 10 * time.Second,
	}

	requestBody := ChatRequest{
		Model: "gpt-3.5-turbo",
		Messages: []Message{
			{Role: "system", Content: GetPrompt()},
			{Role: "user", Content: userInput},
		},
	}

	req, err := buildRequest(requestBody, apiKey, openAIURL)
	if err != nil {
		return "", err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		log.Printf("unexpected status code: %d, body: %s", resp.StatusCode, string(bodyBytes))
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var chatResponse ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&chatResponse); err != nil {
		return "", err
	}

	if len(chatResponse.Choices) == 0 {
		log.Println("no choices in response")
		return "", fmt.Errorf("no choices in response")
	}

	return chatResponse.Choices[0].Message.Content, nil
}

func buildRequest(body ChatRequest, apiKey string, openAIURL string) (*http.Request, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", openAIURL, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	return req, nil
}
