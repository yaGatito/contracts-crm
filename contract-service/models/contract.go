package models

import "time"

type Contract struct {
	ID uint `gorm:"primaryKey"`

	Number    int64 `gorm:"uniqueIndex"`
	StartDate time.Time
	EndDate   time.Time
	Type      string
}
