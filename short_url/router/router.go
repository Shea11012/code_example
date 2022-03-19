package router

import (
	"time"

	"shor_url/handlers"

	"github.com/labstack/echo/v4"
)

func NewServer() *echo.Echo {
	e := echo.New()

	e.GET("/:sid", handlers.GetUrl)
	e.GET("/", func(c echo.Context) error {
		return c.String(200, time.Now().String())
	})

	e.POST("/api/tinyurl", handlers.UrlChange)

	return e
}
