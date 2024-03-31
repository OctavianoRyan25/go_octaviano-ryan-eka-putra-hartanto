package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Fruit struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	Nutrition Nutrition `json:"nutrition"`
}

type Nutrition struct {
	Calories     float64 `json:"calories"`
	Fat          float64 `json:"fat"`
	Sugar        float64 `json:"sugar"`
	Carbohydrate float64 `json:"carbohydrate"`
	Protein      float64 `json:"protein"`
}

var fruits []Fruit

func main() {
	e := echo.New()

	e.GET("/api/v1/fruits", getFruits)
	e.GET("/api/v1/fruits/:id", getFruitByID)
	e.POST("/api/v1/fruits", addFruit)
	e.DELETE("/api/v1/fruits/:id", deleteFruit)

	log.Fatal(e.Start(":8080"))
}

func getFruits(c echo.Context) error {
	return c.JSON(http.StatusOK, fruits)
}

func getFruitByID(c echo.Context) error {
	id := c.Param("id")
	for _, fruit := range fruits {
		if fruit.ID == id {
			return c.JSON(http.StatusOK, fruit)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "Fruit not found"})
}

func addFruit(c echo.Context) error {
	var fruit Fruit
	if err := c.Bind(&fruit); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	fruit.ID = uuid.New().String()

	// GET API dari fruityvice.com
	nutritionURL := fmt.Sprintf("https://www.fruityvice.com/api/fruit/%s", fruit.Name)
	resp, err := http.Get(nutritionURL)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer resp.Body.Close()

	var nutritionData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&nutritionData); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	nutrition, ok := nutritionData["nutritions"].(map[string]interface{})
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch nutrition data"})
	}

	fruit.Nutrition = Nutrition{
		Calories:     nutrition["calories"].(float64),
		Fat:          nutrition["fat"].(float64),
		Sugar:        nutrition["sugar"].(float64),
		Carbohydrate: nutrition["carbohydrates"].(float64),
		Protein:      nutrition["protein"].(float64),
	}

	fruits = append(fruits, fruit)
	return c.JSON(http.StatusCreated, fruit)
}

func deleteFruit(c echo.Context) error {
	id := c.Param("id")
	for i, fruit := range fruits {
		if fruit.ID == id {
			fruits = append(fruits[:i], fruits[i+1:]...)
			return c.JSON(http.StatusOK, map[string]string{"message": "Fruit deleted successfully"})
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "Fruit not found"})
}
