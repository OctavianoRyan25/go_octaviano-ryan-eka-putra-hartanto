package main

import (
	"strings"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Word       Word `json:"word"`
	Length     int  `json:"length"`
	Vocals     int  `json:"num_of_vocals"`
	Palindrome bool `json:"palindrome"`
}

type Word struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func main() {
	e := echo.New()
	e.GET("/word", wordHandler)
	e.POST("/word", wordPostHandler)
	e.Logger.Fatal(e.Start(":8080"))
}

func wordHandler(c echo.Context) error {
	words := generateWord()
	responses := []Response{}
	for _, word := range words {
		response := Response{
			Word:       word,
			Length:     len(word.Message),
			Vocals:     countVocals(word.Message),
			Palindrome: isPalindrome(word.Message),
		}
		responses = append(responses, response)
	}
	return c.JSON(200, responses)
}

func wordPostHandler(c echo.Context) error {
	word := Word{}
	err := c.Bind(&word)
	//word to lower case
	lowerCase := strings.ToLower(word.Message)
	if err != nil {
		return c.JSON(400, err)
	}

	if word.Message == "" {
		errorResponse := ErrorResponse{
			Message: "message is required",
		}
		return c.JSON(400, errorResponse)
	}
	response := Response{
		Word:       word,
		Length:     len(word.Message),
		Vocals:     countVocals(word.Message),
		Palindrome: isPalindrome(lowerCase),
	}
	return c.JSON(200, response)
}

func countVocals(word string) int {
	vocals := "aiueo"
	count := 0
	for _, char := range word {
		if strings.Contains(vocals, string(char)) {
			count++
		}
	}
	return count
}

func isPalindrome(word string) bool {
	reversed := ""
	for i := len(word) - 1; i >= 0; i-- {
		reversed += string(word[i])
	}
	if word != reversed {
		return false
	}
	return true
}

func generateWord() []Word {
	words := []Word{
		{
			Message: "civic",
		},
		{
			Message: "rumah",
		},
	}
	return words
}
