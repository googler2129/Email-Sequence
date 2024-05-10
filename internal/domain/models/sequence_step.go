package models

import (
	"github.com/guregu/null/v5"
	"gorm.io/gorm"
	"time"
)

type SequenceStep struct {
	ID          uint64 `gorm:"primaryKey"`
	Subject     string `gorm:"not null"`
	Content     string `gorm:"not null"`
	WaitingDays uint64 `gorm:"not null"`
	SequenceID  uint64
	SerialOrder uint64
	CreatedAt   time.Time
	CreatedBy   null.String
	CreatedByID null.String
	UpdatedAt   time.Time
	UpdatedBy   null.String
	UpdatedByID null.String
	DeletedAt   gorm.DeletedAt
}
