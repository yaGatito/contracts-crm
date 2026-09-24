package domain

import "time"

type LegalPerson struct {
	Name         string
	ShortName    string
	RegisteredAt time.Time
	LocalKWED    string // ex: 62.01
	LocalERDPOU  int32 // 8-digit
}

type Person struct {
	Lastname    string
	Name        string
	Patronym    string
	DateOfBirth time.Time
	LocalRNOKPP int64 // 10-digit
}
