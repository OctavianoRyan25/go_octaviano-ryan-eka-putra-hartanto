package controllers

import (
	"net/http"
	"prioritas1gorm/configs"
	"prioritas1gorm/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func DeletePackage(c echo.Context) error {
	id := c.Param("id")
	conv, err := strconv.Atoi(id)
	if err != nil {
		BaseResponseError := models.ErrorResponse{
			ErrorCode: 400,
			Status:    "Failed",
			Message:   "Id invalid",
		}
		return c.JSON(http.StatusBadRequest, BaseResponseError)
	}
	err = configs.DB.Where("id = ?", conv).Delete(&models.Package{}).Error
	if err != nil {
		BaseResponseError := models.ErrorResponse{
			ErrorCode: 500,
			Status:    "Failed",
			Message:   err.Error(),
		}
		return c.JSON(http.StatusBadRequest, BaseResponseError)
	}
	SuccessResponse := models.SuccessResponse{
		Status:  "Success",
		Message: "Success delete package",
	}
	return c.JSON(http.StatusOK, SuccessResponse)
}
