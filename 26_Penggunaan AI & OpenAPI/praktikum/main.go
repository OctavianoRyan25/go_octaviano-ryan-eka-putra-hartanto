package main

import (
	"context"
	"fmt"
	"net/http"

	openai "github.com/gtkit/go-openai"
	"github.com/labstack/echo/v4"
)

type Input struct {
	Budget  int    `json:"budget"`
	Purpose string `json:"purpose"`
}

type Response struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

func recommendLaptop(c echo.Context) error {
	// Parse request body
	var input Input
	if err := c.Bind(&input); err != nil {
		return err
	}

	// Initialize OpenAI client
	apiKey := "sk-..."
	if apiKey == "" {
		return c.JSON(http.StatusInternalServerError, Response{Status: "error", Data: "Please set your OpenAI API key."})
	}
	client := openai.NewClient(apiKey)

	// Generate recommendation using OpenAI
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("What is a good %s laptop under Rp %d?", input.Purpose, input.Budget),
				},
			},
		},
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Status: "error", Data: err.Error()})
	}

	// Extract recommendation from response
	recommendation := resp.Choices[0].Message.Content

	// Create response
	response := Response{
		Status: "success",
		Data:   recommendation,
	}

	return c.JSON(http.StatusOK, response)
}

func main() {
	e := echo.New()

	e.POST("/recommend-laptop", recommendLaptop)
	e.Logger.Fatal(e.Start(":8080"))
}
