package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"prioritas1gorm/configs"
	"prioritas1gorm/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func AddFruit(c echo.Context) error {
	fruit := models.Fruit{}
	if err := c.Bind(&fruit); err != nil {
		baseResponseError := models.ErrorResponse{
			ErrorCode: 400,
			Status:    "Failed",
			Message:   "Failed to bind fruit data",
		}
		return c.JSON(http.StatusBadRequest, baseResponseError)
	}

	// konversi id ke uuid
	fruit.ID = uuid.New().String()

	// Fetch data nutrition
	nutritionURL := fmt.Sprintf("https://www.fruityvice.com/api/fruit/%s", fruit.Name)
	resp, err := http.Get(nutritionURL)
	if err != nil {
		baseResponseError := models.ErrorResponse{
			ErrorCode: 500,
			Status:    "Failed",
			Message:   "Failed to fetch nutrition data",
		}
		return c.JSON(http.StatusInternalServerError, baseResponseError)
	}
	defer resp.Body.Close()

	var nutritionData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&nutritionData); err != nil {
		baseResponseError := models.ErrorResponse{
			ErrorCode: 500,
			Status:    "Failed",
			Message:   "Failed to decode nutrition data",
		}
		return c.JSON(http.StatusInternalServerError, baseResponseError)
	}

	nutrition, ok := nutritionData["nutritions"].(map[string]interface{})
	if !ok || nutrition == nil {
		baseResponseError := models.ErrorResponse{
			ErrorCode: 500,
			Status:    "Failed",
			Message:   "Failed to parse nutrition data",
		}
		return c.JSON(http.StatusInternalServerError, baseResponseError)
	}

	fruit.Nutrition = models.Nutrition{
		Calories:     getFloat64Value(nutrition["calories"]),
		Fat:          getFloat64Value(nutrition["fat"]),
		Sugar:        getFloat64Value(nutrition["sugar"]),
		Carbohydrate: getFloat64Value(nutrition["carbohydrate"]),
		Protein:      getFloat64Value(nutrition["protein"]),
	}

	// Simpan data buah beserta nutrisinya ke dalam database
	if err := configs.DB.Create(&fruit).Error; err != nil {
		baseResponseError := models.ErrorResponse{
			ErrorCode: 500,
			Status:    "Failed",
			Message:   "Failed to insert fruit data into database",
		}
		return c.JSON(http.StatusInternalServerError, baseResponseError)
	}

	successResponse := models.SuccessResponse{
		Status:  "Success",
		Message: "Successfully added fruit data",
		Data:    fruit,
	}
	return c.JSON(http.StatusOK, successResponse)
}

func getFloat64Value(value interface{}) float64 {
	if value == nil {
		return 0
	}
	if floatValue, ok := value.(float64); ok {
		return floatValue
	}
	return 0
}
