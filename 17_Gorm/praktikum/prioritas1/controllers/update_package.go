package controllers

import (
	"net/http"
	"prioritas1gorm/configs"
	"prioritas1gorm/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func UpdatePackage(c echo.Context) error {
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
	c.Bind(&Package)

	err = configs.DB.Model(&Package).Where("id = ?", conv).Updates(&Package).Error
	if err != nil {
		BaseResponseError := models.ErrorResponse{
			ErrorCode: 500,
			Status:    "Failed",
			Message:   err.Error(),
		}
		return c.JSON(http.StatusBadRequest, BaseResponseError)
	}

	response := models.PackageResponse{
		Name:             Package.Name,
		Sender:           Package.Sender,
		Receiver:         Package.Receiver,
		SenderLocation:   Package.SenderLocation,
		ReceiverLocation: Package.ReceiverLocation,
		Fee:              Package.Fee,
		Weight:           Package.Weight,
	}

	SuccessResponse := models.SuccessResponse{
		Status:  "Success",
		Message: "Success update package",
		Data:    response,
	}
	return c.JSON(http.StatusOK, SuccessResponse)
}
