package main

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Food struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Responses struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []Food `json:"data"`
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Food   `json:"data"`
}

func main() {
	// Code
	e := echo.New()

	//ROUTE
	e.GET("/api/v1/foods", getFoods)
	e.GET("/api/v1/foods/:id", getFoodsById)
	e.POST("/api/v1/foods", insertFood)
	e.PUT("/api/v1/foods/:id", updateFood)
	e.DELETE("/api/v1/foods/:id", deleteFood)
	e.Logger.Fatal(e.Start(":8080"))
}

func getFoods(c echo.Context) error {
	foods := generateFoods()
	return c.JSON(200, Responses{
		Code:    200,
		Message: "Success",
		Data:    foods,
	})
}

func getFoodsById(c echo.Context) error {
	id := c.Param("id")
	parse, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error", err)
		return c.JSON(400, ErrorResponse{
			Code:    400,
			Message: "Bad Request",
		})
	}
	foods := generateFoods()
	var food Food
	found := false
	for _, f := range foods {
		if f.Id == parse {
			food = f
			found = true
			break
		}
	}
	if !found {
		return c.JSON(404, ErrorResponse{
			Code:    404,
			Message: "Data Not Found",
		})
	}
	return c.JSON(200, Response{
		Code:    200,
		Message: "Success",
		Data:    food,
	})
}

func insertFood(c echo.Context) error {
	food := Food{}
	err := c.Bind(&food)
	if err != nil {
		fmt.Println("Error", err)
		return c.JSON(400, ErrorResponse{
			Code:    400,
			Message: err.Error(),
		})
	}
	return c.JSON(200, Response{
		Code:    200,
		Message: "Success",
		Data:    food,
	})
}

func updateFood(c echo.Context) error {
	id := c.Param("id")
	parse, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error", err)
		return c.JSON(400, ErrorResponse{
			Code:    400,
			Message: "Bad Request",
		})
	}
	food := Food{}
	food.Id = parse
	found := false
	for _, f := range generateFoods() {
		if f.Id == parse {
			food = f
			found = true
			break
		}
	}
	if !found {
		return c.JSON(404, ErrorResponse{
			Code:    404,
			Message: "Data Not Found",
		})
	}
	err = c.Bind(&food)
	if err != nil {
		return nil
	}
	return c.JSON(200, Response{
		Code:    200,
		Message: "Success",
		Data:    food,
	})
}

func deleteFood(c echo.Context) error {
	id := c.Param("id")
	parse, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error", err)
		return c.JSON(400, ErrorResponse{
			Code:    400,
			Message: "Bad Request",
		})
	}
	foods := generateFoods()
	var food Food
	found := false
	for _, f := range foods {
		if f.Id == parse {
			food = f
			found = true
			break
		}
	}
	if !found {
		return c.JSON(404, ErrorResponse{
			Code:    404,
			Message: "Data Not Found",
		})
	}
	return c.JSON(200, Response{
		Code:    200,
		Message: "Success",
		Data:    food,
	})
}

func generateFoods() []Food {
	foods := []Food{
		{
			Id:          1,
			Name:        "Nasi Goreng",
			Price:       12000,
			Description: "Delicious",
		},
		{
			Id:          2,
			Name:        "Ayam Goreng",
			Price:       22000,
			Description: "Delicious",
		},
	}
	return foods
}
