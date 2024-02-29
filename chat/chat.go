package chat

import (
	"context"
	"github.com/sashabaranov/go-openai"
)

type BaseChat interface {
	Chat(userID string, msg string) string
}

type ErrorChat struct {
	errMsg string
}

func (e *ErrorChat) Chat(userID string, msg string) string {
	return e.errMsg
}

type Echo struct{}

func (e *Echo) Chat(userID string, msg string) string {
	return msg
}

type SimpleGptChat struct {
	token string
	url   string
}

func (s *SimpleGptChat) Chat(userID string, msg string) string {
	cfg := openai.DefaultConfig(s.token)
	cfg.BaseURL = s.url
	client := openai.NewClientWithConfig(cfg)
	resp, err := client.CreateChatCompletion(context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: msg,
				},
			},
		})
	if err != nil {
		return err.Error()
	}
	return resp.Choices[0].Message.Content
}

func GetChatBot() BaseChat {
	return &Echo{}
}
