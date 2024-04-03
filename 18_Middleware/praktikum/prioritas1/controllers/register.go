package controllers

import (
	"prioritas1gorm/configs"
	"prioritas1gorm/helper"
	"prioritas1gorm/models"

	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	user.Password = helper.HashPass(user.Password)

	err := configs.DB.Create(&user).Error

	if err != nil {
		BaseResponseError := models.ErrorResponse{
			ErrorCode: 400,
			Status:    "Failed",
			Message:   "Failed to register",
		}
		return c.JSON(400, BaseResponseError)
	}

	BaseResponseSuccess := models.SuccessResponse{
		Status:  "Success",
		Message: "Register successfully",
		Data:    user,
	}

	return c.JSON(200, BaseResponseSuccess)
}
