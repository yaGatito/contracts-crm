package models

import "time"

type LegalPerson struct {
	ID uint `gorm:"primaryKey"`

	Name         string
	ShortName    string
	RegisteredAt time.Time
	LocalKWED    string `gorm:"column:local_kwed"` // ex: 62.01
	LocalEDRPOU  uint   `gorm:"uniqueIndex"`       // 8-digit
}

type Person struct {
	ID uint `gorm:"primaryKey"`

	Lastname    string
	Name        string
	Patronym    string
	DateOfBirth time.Time
	LocalRNOKPP uint64 `gorm:"uniqueIndex"` // 10-digit
}
