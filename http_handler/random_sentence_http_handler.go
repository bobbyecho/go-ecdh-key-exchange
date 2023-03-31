package http_handler

import (
	"github.com/labstack/echo"
	"key-exchange/utils"
	"math/rand"
	"net/http"
)

var randomSentence []string

func init() {
	randomSentence = []string{
		"Jenny made the announcement that her baby was an alien.",
		"This made him feel like an old-style rootbeer float smells.",
		"25 years later, she still regretted that specific moment.",
		"There should have been a time and a place, but this wasn't it.",
		"Her life in the confines of the house became her new normal.",
		"The dead trees waited to be ignited by the smallest spark and seek their revenge.",
	}
}

type ResponseSentence struct {
	Sentence string `json:"sentence"`
}

func HandleRandomSentence(c echo.Context) error {
	sentenceLength := len(randomSentence) - 1
	randomIdx := rand.Intn(sentenceLength-0) + 0
	sentence := randomSentence[randomIdx]

	encryptedSentence := utils.EncryptPayload(c.Request().Header.Get("SessionID"), sentence)

	return c.JSON(http.StatusOK, ResponseSentence{
		Sentence: string(encryptedSentence),
	})
}
