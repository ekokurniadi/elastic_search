package repository

import (
	"elastic_go/entity"

	"gorm.io/gorm"
)

type TweetRepository interface {
	Save(tweet entity.Tweet) (entity.Tweet, error)
}

type tweetRepository struct {
	db *gorm.DB
}

func NewTweetRepository(db *gorm.DB) *tweetRepository {
	return &tweetRepository{db}
}

func (r *tweetRepository) Save(tweet entity.Tweet) (entity.Tweet, error) {
	err := r.db.Create(&tweet).Error
	if err != nil {
		return tweet, err
	}
	return tweet, nil
}
