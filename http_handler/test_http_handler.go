package http_handler

import (
	"github.com/labstack/echo"
	"net/http"
)

func HandleTest(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}
