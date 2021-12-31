package handler

import (
	"elastic_go/service"

	"github.com/gin-gonic/gin"
)

type tweetHandler struct {
	service service.TweetService
}

func NewTweetHandler(service service.TweetService) *tweetHandler {
	return &tweetHandler{service}
}

func (h *tweetHandler) Save(c *gin.Context) {

}
