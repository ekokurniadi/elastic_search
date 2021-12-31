package entity

import "time"

type Tweet struct {
	User     string    `json:"user"`
	Message  string    `json:"message"`
	Retweets int       `json:"retweets"`
	Created  time.Time `json:"created"`
	Attrs    string    `json:"attributes,omitempty"`
}
