package es

import (
	"context"
	"elastic_go/input"
	"fmt"
	"strconv"

	"github.com/olivere/elastic"
)

type EsService interface {
	Save(ID int, inputData input.TweetInput)
}

type esService struct {
	client *elastic.Client
}

func NewEsService(client *elastic.Client) *esService {
	return &esService{client}
}

func (els *esService) Save(ID int, inputData input.TweetInput) {
	id := strconv.Itoa(ID)
	data, _ := els.client.Index().
		Index("belajar2").
		Id(id).
		BodyJson(&inputData).
		Refresh("true").
		Do(context.TODO())
	fmt.Println(data)

}
