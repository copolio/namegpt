package chatgpt

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/copolio/namegpt/config"
	"github.com/sashabaranov/go-openai"
	"log"
)

var client *openai.Client

func init() {
	curConfig := config.Get()
	log.Default().Println(curConfig.ChatgptToken)
	client = openai.NewClient(curConfig.ChatgptToken)
}

func GetSimilarDomains(keyword string) (domains []string, err error) {
	response, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are an assistant that only speaks JSON array. Do not write normal text. Give 50 similar domain names without tld given input",
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
