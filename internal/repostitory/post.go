package repostitory

import "time"

type Post struct {
	title     string
	text      string
	tag       string
	createdAt time.Time
}

func NewPost(title string, text string, tag string) *Post {
	return &Post{
		title:     title,
		text:      text,
		tag:       tag,
		createdAt: time.Now(),
	}
}
