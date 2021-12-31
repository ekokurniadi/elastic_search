package input

import "time"

type TweetInput struct {
	User     string     `json:"user"`
	Message  string     `json:"message"`
	Retweets int        `json:"retweets"`
	Created  time.Time  `json:"created"`
	Attrs    Attributes `json:"attributes,omitempty"`
}

type Attributes struct {
	Views int  `json:"views"`
	Vip   bool `json:"vip"`
}
