package controllers

import (
	"errors"
	"net/http"
	"prioritas1gorm/configs"
	"prioritas1gorm/models"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func FindPackageById(c echo.Context) error {
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
	Package := models.Package{}
	err = configs.DB.Find(&Package, conv).Error

	//Error handling
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Paket tidak ditemukan
			BaseResponseError := models.ErrorResponse{
				ErrorCode: 404,
				Status:    "Failed",
				Message:   "Package not found",
			}
			return c.JSON(http.StatusNotFound, BaseResponseError)
		}
		// Terjadi kesalahan lain dalam database
		BaseResponseError := models.ErrorResponse{
			ErrorCode: 500,
			Status:    "Failed",
			Message:   "Internal server error",
		}
		return c.JSON(http.StatusInternalServerError, BaseResponseError)
	}

	if Package.ID == 0 {
		// Paket tidak ditemukan
		BaseResponseError := models.ErrorResponse{
			ErrorCode: 404,
			Status:    "Failed",
			Message:   "Package not found",
		}
		return c.JSON(http.StatusNotFound, BaseResponseError)

	}

	// Jika berhasil, kembalikan respons sukses
	SuccessResponse := models.SuccessResponse{
		Status:  "Success",
		Message: "Success find package",
		Data:    Package,
	}
	return c.JSON(http.StatusOK, SuccessResponse)
}
