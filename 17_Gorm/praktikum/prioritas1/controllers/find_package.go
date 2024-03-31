package controllers

import (
	"net/http"
	"prioritas1gorm/configs"
	"prioritas1gorm/models"

	"github.com/labstack/echo/v4"
)

func FindPackage(c echo.Context) error {
	packages := []models.Package{}
	result := configs.DB.Find(&packages)
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
		Data:    packages,
	}
	return c.JSON(http.StatusOK, SuccessResponse)
}
