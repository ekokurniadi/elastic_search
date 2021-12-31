package service

import (
	"elastic_go/entity"
	"elastic_go/input"
	"elastic_go/repository"
)

type TweetService interface {
	Save(input input.TweetInput) (entity.Tweet, error)
}

type tweetService struct {
	tweetRepository repository.TweetRepository
}

func NewTweetService(tweetRepository repository.TweetRepository) *tweetService {
	return &tweetService{tweetRepository}
}

func (s *tweetService) Save(input input.TweetInput) (entity.Tweet, error) {

}
