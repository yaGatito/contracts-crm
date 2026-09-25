package models

import (
	"time"
)

type Contract struct {
	ID uint `gorm:"primaryKey"`

	Number    uint64 `gorm:"uniqueIndex"`
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

type ContractDetails struct {
	Contract     Contract
	Persons      []Person
	LegalPersons []LegalPerson
}

type SearchPersonContractsFilter struct {
	ContractType  string
	StartDateFrom time.Time
	StartDateTo   time.Time

	PersonName        string
	PersonLastname    string
	PersonPatronym    string
	PersonDateOfBirth time.Time
}

type SearchLegalPersonContractsFilter struct {
	ContractType  string
	StartDateFrom time.Time
	StartDateTo   time.Time

	LegalPersonShortName string
	LegalPersonName      string
	LocalEDRPOU          uint32
}
