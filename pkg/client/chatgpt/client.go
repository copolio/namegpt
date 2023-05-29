package chatgpt

import (
	"context"
	"github.com/copolio/namegpt/config"
	"github.com/sashabaranov/go-openai"
	"log"
)

var client *openai.Client

func init() {
	curConfig := config.Get()
	client = openai.NewClient(curConfig.ChatgptToken)
}

func AskAsSystem(query string) (openai.ChatCompletionResponse, error) {
	response, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: query,
				},
			},
		},
	)

	if err != nil {
		log.Default().Printf("ChatCompletion Error: %v\n", err)
	}
	return response, err
}

func AskAsAssistant(query string) (openai.ChatCompletionResponse, error) {
	response, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: query,
				},
			},
		},
	)

	if err != nil {
		log.Default().Printf("ChatCompletion Error: %v\n", err)
	}
	return response, err
}

func AskAsUser(query string) (openai.ChatCompletionResponse, error) {
	response, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: query,
				},
			},
		},
	)

	if err != nil {
		log.Default().Printf("ChatCompletion Error: %v\n", err)
	}
	return response, err
}
