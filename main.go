package main

import (
	"go-react-example/web"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	web.RegisterHandlers(e)
	e.GET("/api", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
