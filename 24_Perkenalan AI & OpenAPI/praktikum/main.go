package main

import (
	"context"
	"fmt"

	openai "github.com/gtkit/go-openai"
)

func main() {
	apiKey := "sk-..."
	if apiKey == "" {
		fmt.Println("Please set your OpenAI API key.")
		return
	}

	client := openai.NewClient(apiKey)

	fmt.Print("Dimana ibukota provinsi Jawa Tengah? ")

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Dimana ibukota provinsi Jawa Tengah?",
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
