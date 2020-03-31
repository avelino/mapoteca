package models

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	ID        uuid.UUID `gorm:"type:uuid;PRIMARY_KEY;unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string `gorm:"size:255;not null"`
	Subtitle  string `gorm:"size:255;"`
	Slug      string `gorm:"size:255;unique;not null"`
	ImagePath string `gorm:"size:255"`
	Category  string `gorm:"size:16;not null"`
	Content   string `gorm:"type:text;not null"`
}
