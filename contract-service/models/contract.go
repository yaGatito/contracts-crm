package models

import (
	"time"
)

type Contract struct {
	ID uint `gorm:"primaryKey"`

	Number    int64 `gorm:"uniqueIndex"`
	StartDate time.Time
	EndDate   time.Time
	Type      string
}

type ContractLegalPerson struct {
	ContractID    uint `gorm:"primaryKey"`
	LegalPersonID uint `gorm:"primaryKey"`

	LegalPerson LegalPerson `gorm:"foreignKey:LegalPersonID"`
}

type ContractPerson struct {
	ContractID uint `gorm:"primaryKey"`
	PersonID   uint `gorm:"primaryKey"`

	Person Person `gorm:"foreignKey:PersonID"`
}
