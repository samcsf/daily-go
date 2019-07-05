package model

import (
	"time"
)

type Tag struct {
	Id   string `json:id`
	Name string `json:name`
}

type Post struct {
	Id          string    `json:id`
	Title       string    `json:title`
	Tag         []*Tag    `json:tag`
	Content     string    `json:content`
	Month       int       `json:month`
	Year        int       `json:year`
	Create_at   time.Time `json:create_at`
	Modified_at time.Time `json:modified_at`
}
