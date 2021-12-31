package handler

import (
	"elastic_go/input"
	"elastic_go/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type tweetHandler struct {
	service service.TweetService
}

func NewTweetHandler(service service.TweetService) *tweetHandler {
	return &tweetHandler{service}
}

func (h *tweetHandler) Save(c *gin.Context) {
	var input input.TweetInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Gagal mendapatkan data")
		return
	}
	fmt.Println(input)
}
