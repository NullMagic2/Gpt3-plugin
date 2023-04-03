package main

import (
	"context"
	"errors"
	openai "github.com/sashabaranov/go-openai"
	"github.com/ian-kent/gptchat/module"
)

var Plugin module.Plugin = GPT3Plugin{}

type GPT3Plugin struct{}

func (c GPT3Plugin) ID() string {
	return "gpt3"
}

func (c GPT3Plugin) Example() string {
	return "/gpt3 {\"message\": \"What is the capital of France?\"}"
}

func (c GPT3Plugin) Execute(input map[string]interface{}) (map[string]interface{}, error) {
	client := openai.NewClient("your token")

	message, ok := input["message"].(string)
	if !ok {
		return nil, errors.New("Message not provided")
	}

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.Davinci,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: message,
				},
			},
		},
	)

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"response": resp.Choices[0].Message.Content,
	}, nil
}