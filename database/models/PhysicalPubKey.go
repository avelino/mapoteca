package models

import (
	"github.com/google/uuid"
	"time"
)

type PhysicalPubKey struct {
	ID            uuid.UUID `gorm:"type:uuid;PRIMARY_KEY;unique;not null"`
	KeyIdentifier string    `gorm:"size:12;unique;not null"`
	AuthTokens    []AuthToken
	CreatedAt     time.Time
}
