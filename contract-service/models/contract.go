package models

import "time"

type Contract struct {
	ID uint `gorm:"primaryKey"`

	Number    int64 `gorm:"uniqueIndex"`
	StartDate time.Time
	EndDate   time.Time
	Type      string
}

type ContractLegalPerson struct {
	ID uint `gorm:"primaryKey"`

	ContractID    uint `gorm:"not null"`
	LegalPersonID uint `gorm:"not null"`
}

type ContractPerson struct {
	ID uint `gorm:"primaryKey"`

	ContractID uint `gorm:"not null"`
	PersonID   uint `gorm:"not null"`
}
