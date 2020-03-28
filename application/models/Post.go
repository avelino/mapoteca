package models

import (
	"time"
)

type Post struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Subtitle  string    `json:"subtitle"`
	Slug      string    `json:"slug"`
	ImagePath string    `json:"imagePath"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"createdAt"`
}
