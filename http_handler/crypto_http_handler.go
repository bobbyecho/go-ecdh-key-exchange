package http_handler

import (
	"encoding/hex"
	"github.com/labstack/echo"
	"key-exchange/crypto"
	"net/http"
)

type keyExchangeRequest struct {
	PublicKey string `json:"publicKey"`
}

type keyExchangeResponse struct {
	SessionID string `json:"sessionId"`
	PublicKey string `json:"publicKey"`
}

func HandleKeyExchange(c echo.Context) error {
	reqData := new(keyExchangeRequest)
	if err := c.Bind(reqData); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if reqData.PublicKey == "" {
		return c.JSON(http.StatusUnprocessableEntity, "`publicKey is required`")
	}

	sessionID, pubKeyString, err := crypto.EcdhKeyExchange([]byte(reqData.PublicKey))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, keyExchangeResponse{
		SessionID: sessionID,
		PublicKey: hex.EncodeToString(pubKeyString),
	})
}
