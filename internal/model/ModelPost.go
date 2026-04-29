package model

import "time"

type PostModel struct {
	Title     string
	Text      string
	Tag       string
	CreatedAt time.Time
}

type EncodePostModel struct {
	Text string `json:"content"`
	Tag  string `json:"tag"`
}

type ErrorModel struct {
	error string
	data  time.Time
}

func NewErrorModel(err error) ErrorModel {
	return ErrorModel{error: err.Error(), data: time.Now()}
}

func ValidateEncodePostModel(post PostModel) EncodePostModel {
	return EncodePostModel{Text: post.Text, Tag: post.Tag}
}
