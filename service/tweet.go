package service

import (
	"elastic_go/entity"
	"elastic_go/input"
	"elastic_go/repository"
	"encoding/json"
	"time"
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
	tweet := entity.Tweet{}
	tweet.User = input.User
	tweet.Message = input.Message
	tweet.Retweets = input.Retweets
	tweet.Created = time.Now()

	x := make(map[string]interface{})
	for key, val := range input.Attrs {
		x[key] = val
	}

	jsonByte, _ := json.Marshal(x)
	tweet.Attrs = string(jsonByte)

	newTweet, err := s.tweetRepository.Save(tweet)
	if err != nil {
		return newTweet, err
	}

	return newTweet, nil
}
