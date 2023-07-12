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

func GenerateDomainNames(keyword string) (domains []string, err error) {
	response, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleSystem,
					Content: "Given the input below, create JSON array containing 10 domain names as string without tld. " +
						"If input is domain, give similar domain. Else if it is description, recommend domain names. " +
						"Do not write normal text. " +
						"Do not write normal text. " +
						"Do not repeat last character or append numbers. " +
						"Domain names should be in English, preferably without hyphen. ",
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

func GetRecommendedDomainStream(keyword string) (*openai.ChatCompletionStream, error) {
	completionStream, err := client.CreateChatCompletionStream(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "Given the description below, create JSON array containing 10 recommended domain names without tld. Do not write normal text. Do not repeat last character or append numbers",
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
