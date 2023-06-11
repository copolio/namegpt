package chatgpt

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/copolio/namegpt/config"
	"github.com/sashabaranov/go-openai"
)

var client *openai.Client

func init() {
	client = openai.NewClient(config.NameGptAppConfig.Chatgpt.Token)
}

func GetSimilarDomains(keyword string) (domains []string, err error) {
	response, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "Given the input below, create JSON array containing 50 similar domain names without tld. Do not write normal text.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: keyword,
				},
			},
		},
	)

	if err != nil {
		return nil, fmt.Errorf("ChatGPT completion failed: %w", err)
	}

	message := response.Choices[0].Message.Content
	err = json.Unmarshal([]byte(message), &domains)
	if err != nil {
		return nil, fmt.Errorf("JSON parse failed (message: %s): %w", message, err)
	}

	return domains, err
}

func GetSimilarDomainStream(keyword string) (*openai.ChatCompletionStream, error) {
	completionStream, err := client.CreateChatCompletionStream(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "Given the input below, create JSON array containing 50 similar, unregistered domain names without tld. Do not write normal text.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: keyword,
				},
			},
			Stream: true,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("ChatGPT Stream failed: %w", err)
	}

	return completionStream, err
}
