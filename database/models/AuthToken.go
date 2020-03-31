package models

import (
	"github.com/google/uuid"
	"time"
)

type AuthToken struct {
	ID              uuid.UUID `gorm:"type:uuid;PRIMARY_KEY;unique;not null"`
	KeyIdentifierId uuid.UUID `gorm:"type:uuid"`
	CreatedAt       time.Time
}
