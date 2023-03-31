package http_handler

import (
	"github.com/labstack/echo"
	"key-exchange/utils"
	"net/http"
)

type ResponseSomething struct {
	Status      bool   `json:"status"`
	Translation string `json:"translation"`
}

type RequestSomething struct {
	Payload string `json:"payload"`
}

func HandleSendSomething(c echo.Context) error {
	request := new(RequestSomething)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ResponseSomething{
			Status:      false,
			Translation: "",
		})
	}

	decryptPayload := utils.DecryptPayload(c.Request().Header.Get("SessionID"), request.Payload)

	return c.JSON(http.StatusOK, ResponseSomething{
		Status:      true,
		Translation: string(decryptPayload),
	})
}
