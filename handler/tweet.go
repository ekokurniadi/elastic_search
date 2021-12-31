package handler

import (
	"context"
	"elastic_go/formatter"
	"elastic_go/input"
	"elastic_go/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
)

type tweetHandler struct {
	service service.TweetService
	client  *elastic.Client
}

func NewTweetHandler(service service.TweetService, client *elastic.Client) *tweetHandler {
	return &tweetHandler{service, client}
}

func (h *tweetHandler) Save(c *gin.Context) {
	var input input.TweetInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Gagal mendapatkan data")
		return
	}
	newTweet, err := h.service.Save(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Gagal mendapatkan data")
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	_, _ = h.EsService(newTweet.ID, input)

	c.JSON(http.StatusOK, formatter.FormatTweet(newTweet))
}

func (h *tweetHandler) EsService(ID int, input input.TweetInput) (*elastic.IndexResponse, error) {

	id := strconv.Itoa(ID)
	input.Created = time.Now()
	data, err := h.client.Index().Index("belajar2").Id(id).Type("_doc").BodyJson(&input).Refresh("true").Do(context.TODO())
	if err != nil {
		return data, err
	}
	return data, nil
}
