package formatter

import (
	"elastic_go/entity"
	"encoding/json"
	"time"
)

type TweetFormatter struct {
	ID       int                    `json:"id"`
	User     string                 `json:"user"`
	Message  string                 `json:"message"`
	Retweets int                    `json:"retweets"`
	Created  time.Time              `json:"created"`
	Attrs    map[string]interface{} `json:"attributes"`
}

func FormatTweet(tweet entity.Tweet) TweetFormatter {
	tweetFormatter := TweetFormatter{}
	tweetFormatter.ID = tweet.ID
	tweetFormatter.User = tweet.User
	tweetFormatter.Retweets = tweet.Retweets
	tweetFormatter.Message = tweet.Message
	tweetFormatter.Created = tweet.Created

	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(tweet.Attrs), &jsonMap)

	tweetFormatter.Attrs = jsonMap
	return tweetFormatter
}
