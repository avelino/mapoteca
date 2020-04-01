package models

import (
	"github.com/google/uuid"
)

type HttpSuccess struct {
	Greetings    string      `json:"greetings"`
	GeneratedIds []uuid.UUID `json:"generated_ids"`
}
