package controllers

import (
	"net/http"
	"prioritas1gorm/configs"
	"prioritas1gorm/models"

	"github.com/labstack/echo/v4"
)

func AddPackage(c echo.Context) error {
	Package := models.Package{}
	c.Bind(&Package)

	result := configs.DB.Create(&Package)

	if result.Error != nil {
		BaseResponseError := models.ErrorResponse{
			ErrorCode: 500,
			Status:    "Failed",
			Message:   "Failed to add insert database`",
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
		Message: "Success add package",
		Data:    response,
	}
	return c.JSON(http.StatusOK, SuccessResponse)
}
