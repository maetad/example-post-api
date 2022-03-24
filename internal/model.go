package internal

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title    string    `json:"title"`
	Detail   string    `json:"detail"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	gorm.Model
	PostID  uint   `json:"-"`
	Message string `json:"message"`
}
