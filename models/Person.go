package models

import (
	"github.com/google/uuid"
	"time"
)

type Person struct {
	ID        uuid.UUID `gorm:"type:uuid;PRIMARY_KEY;unique;not null" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Email     string    `gorm:"size:255;not null" json:"email"`
	FirstName string    `gorm:"size:255;not null" json:"firstName"`
	LastName  string    `gorm:"size:255;not null" json:"lastName"`
	Posts     []Post    `gorm:"foreignKey:PersonID"`
}
