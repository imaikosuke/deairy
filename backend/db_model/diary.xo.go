package db_model

// Code generated by xo. DO NOT EDIT.

import (
	"time"
)

// Diary represents a row from 'deairy.diaries'.
type Diary struct {
	ID         string    `json:"id"`         // id
	UserID     string    `json:"user_id"`    // user_id
	Title      string    `json:"title"`      // title
	Content    string    `json:"content"`    // content
	Visibility string    `json:"visibility"` // visibility
	CreatedAt  time.Time `json:"created_at"` // created_at
	UpdatedAt  time.Time `json:"updated_at"` // updated_at
}