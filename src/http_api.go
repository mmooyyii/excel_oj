package src

import (
	"github.com/labstack/echo"
	"net/http"
)

func HttpServer() *echo.Echo {
	e := echo.New()
	e.GET("/", Index)
	return e
}

func Index(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}
