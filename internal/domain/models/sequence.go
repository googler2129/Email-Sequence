package models

import (
	"github.com/guregu/null/v5"
	"gorm.io/gorm"
	"time"
)

type Sequence struct {
	ID                   uint64 `gorm:"primaryKey"`
	Name                 string `gorm:"not null"`
	OpenTrackingEnabled  bool   `gorm:"not null"`
	ClickTrackingEnabled bool   `gorm:"not null"`
	SequenceSteps        []*SequenceStep
	CreatedAt            time.Time
	CreatedBy            null.String
	CreatedByID          null.String
	UpdatedAt            time.Time
	UpdatedBy            null.String
	UpdatedByID          null.String
	DeletedAt            gorm.DeletedAt
}
