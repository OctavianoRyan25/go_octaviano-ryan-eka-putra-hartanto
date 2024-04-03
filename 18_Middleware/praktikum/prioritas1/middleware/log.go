package middleware

import (
	"github.com/labstack/echo/v4"

	"github.com/labstack/echo/v4/middleware"
)

type LogConfig struct {
	Format string
}

func (c *LogConfig) Init() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: c.Format,
	})
}
