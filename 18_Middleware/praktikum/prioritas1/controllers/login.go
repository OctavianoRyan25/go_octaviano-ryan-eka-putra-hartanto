package controllers

import (
	"net/http"
	"prioritas1gorm/configs"
	"prioritas1gorm/helper"
	"prioritas1gorm/models"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	user := models.User{}
	password := ""
	c.Bind(&user)

	password = user.Password

	err := configs.DB.Where("email = ?", user.Email).First(&user).Error
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error":   "Unauthorized",
			"message": "Invalid email/password",
		})
	}

	comparePass := helper.ComparePass([]byte(user.Password), []byte(password))
	if !comparePass {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error":   "Unauthorized",
			"message": "Invalid email/password",
		})
	}

	token, err := helper.GenerateToken(user.ID, user.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "Login successfully",
		"ID":     user.ID,
		"token":  token,
	})
}
