package api

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
}

func NewClient(apiKey, baseURL string) *Client {
	if baseURL == "" {
		baseURL = "https://router.huggingface.co/v1/chat/completions"
	}
	return &Client{
		APIKey:  apiKey,
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
}

type ChatRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Stream      bool      `json:"stream,omitempty"`
	MaxTokens   int       `json:"max_tokens,omitempty"`
	Temperature float64   `json:"temperature,omitempty"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatResponse struct {
	Choices []struct {
		Index   int     `json:"index"`
		Message Message `json:"message"`
	} `json:"choices"`
}

type StreamingChatResponse struct {
	Choices []struct {
		Index int `json:"index"`
		Delta struct {
			Content string `json:"content"`
		} `json:"delta"`
		FinishReason string `json:"finish_reason,omitempty"`
	} `json:"choices"`
}

type StreamingCallback func(chunk string, done bool, err error)

func (c *Client) buildPayload(req ChatRequest) map[string]interface{} {
	payload := map[string]interface{}{
		"model":    req.Model,
		"messages": req.Messages,
		"stream":   req.Stream,
	}

	if req.MaxTokens > 0 {
		payload["max_tokens"] = req.MaxTokens
	}
	if req.Temperature > 0 {
		payload["temperature"] = req.Temperature
	}

	return payload
}

func (c *Client) makeRequest(payload map[string]interface{}) (*http.Response, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequest("POST", c.BaseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+c.APIKey)

	resp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		errorBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode, string(errorBody))
	}

	return resp, nil
}

func (c *Client) SendChat(req ChatRequest) (*ChatResponse, error) {
	req.Stream = false
	payload := c.buildPayload(req)

	resp, err := c.makeRequest(payload)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var chatResp ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&chatResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &chatResp, nil
}

func (c *Client) SendChatStream(req ChatRequest, callback StreamingCallback) error {
	req.Stream = true
	payload := c.buildPayload(req)

	resp, err := c.makeRequest(payload)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("failed to read streaming response: %w", err)
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if line == "data: [DONE]" {
			callback("", true, nil)
			break
		}

		if strings.HasPrefix(line, "data: ") {
			data := strings.TrimPrefix(line, "data: ")
			var streamResp StreamingChatResponse
			if err := json.Unmarshal([]byte(data), &streamResp); err != nil {
				continue
			}

			if len(streamResp.Choices) > 0 {
				choice := streamResp.Choices[0]
				if choice.Delta.Content != "" {
					callback(choice.Delta.Content, false, nil)
				}
				if choice.FinishReason != "" {
					callback("", true, nil)
					break
				}
			}
		}
	}

	return nil
}

func (c *Client) SendChatStreamWithChannel(req ChatRequest) (<-chan string, <-chan error) {
	chunkChan := make(chan string, 100)
	errChan := make(chan error, 1)

	go func() {
		defer func() {
			close(chunkChan)
			close(errChan)
		}()

		err := c.SendChatStream(req, func(chunk string, done bool, err error) {
			if err != nil {
				errChan <- err
				return
			}
			if done {
				return
			}
			if chunk != "" {
				chunkChan <- chunk
			}
		})

		if err != nil {
			errChan <- err
		}
	}()

	return chunkChan, errChan
}
