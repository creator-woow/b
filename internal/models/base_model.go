package models

import (
	"time"
)

// ID declares type of id field for all models
type ID uint64

type BaseModel struct {
	ID        ID        `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
