package controllers

import (
	"net/http"
	"prioritas1gorm/configs"
	"prioritas1gorm/models"

	"github.com/labstack/echo/v4"
)

func FindFruit(c echo.Context) error {
	fruits := []models.Fruit{}
	result := configs.DB.Preload("Nutrition").Find(&fruits)
	if result.Error != nil {
		BaseResponseError := models.ErrorResponse{
			ErrorCode: 500,
			Status:    "Failed",
			Message:   "Failed to add insert database`",
		}
		return c.JSON(http.StatusBadRequest, BaseResponseError)
	}
	SuccessResponse := models.SuccessResponse{
		Status:  "Success",
		Message: "Success find package",
		Data:    fruits,
	}
	return c.JSON(http.StatusOK, SuccessResponse)
}
