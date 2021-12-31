package service

import (
	"elastic_go/entity"
	"elastic_go/input"
	"elastic_go/repository"
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
	tweet.Attrs = input.Attrs
	newTweet, err := s.tweetRepository.Save(tweet)
	if err != nil {
		return newTweet, err
	}
	return newTweet, nil
}
