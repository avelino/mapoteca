package models

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	ID        uuid.UUID `gorm:"type:uuid;PRIMARY_KEY;unique;not null" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Title     string    `gorm:"size:255;not null" json:"title"`
	Subtitle  string    `gorm:"size:255;" json:"subtitle"`
	Slug      string    `gorm:"size:255;unique;not null" json:"slug"`
	ImagePath string    `gorm:"size:255" json:"imagePath"`
	Category  string    `gorm:"size:16;not null" json:"category"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	PersonID  uuid.UUID `json:"personId"`
}
