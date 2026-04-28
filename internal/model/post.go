package model

import "time"

type PostModel struct {
	Title     string
	Text      string
	Tag       string
	CreatedAt time.Time
}

type ErrorModel struct {
	error string
	data  time.Time
}

func NewErrorModel(err error) ErrorModel {
	return ErrorModel{error: err.Error(), data: time.Now()}
}
