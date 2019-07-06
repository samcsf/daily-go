package model

import (
	"time"
)

type Tag struct {
	Id   string `json:id`
	Name string `json:name`
}

type Post struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Create_at   time.Time `json:"create_at"`
	Modified_at time.Time `json:"modified_at"`
}
