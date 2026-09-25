package httpadapter

import (
	"contract-service/models"
	"time"
)

type Contract struct {
	ID        uint      `json:"id,omitempty"`
	Number    int64     `json:"number"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Type      string    `json:"type"`
}

type LegalPerson struct {
	ID           uint      `json:"id,omitempty"`
	Name         string    `json:"name"`
	ShortName    string    `json:"shortname"`
	RegisteredAt time.Time `json:"registered_at"`
	LocalKWED    string    `json:"local_kwed"`
	LocalEDRPOU  int32     `json:"local_edrpou"`
}

type Person struct {
	ID          uint      `json:"id,omitempty"`
	Lastname    string    `json:"lastname"`
	Name        string    `json:"name"`
	Patronym    string    `json:"patronym"`
	DateOfBirth time.Time `json:"date_of_birth"`
	LocalRNOKPP int64     `json:"local_rnokpp"`
}

func personToDto(p models.Person) Person {
	return Person{
		ID:          p.ID,
		Lastname:    p.Lastname,
		Name:        p.Name,
		Patronym:    p.Patronym,
		DateOfBirth: p.DateOfBirth,
		LocalRNOKPP: p.LocalRNOKPP,
	}
}

func legalPersonToDto(lp models.LegalPerson) LegalPerson {
	return LegalPerson{
		ID:           lp.ID,
		Name:         lp.Name,
		ShortName:    lp.ShortName,
		RegisteredAt: lp.RegisteredAt,
		LocalKWED:    lp.LocalKWED,
		LocalEDRPOU:  lp.LocalEDRPOU,
	}
}

func contractToDto(c models.Contract) Contract {
	return Contract{
		ID:        c.ID,
		Number:    c.Number,
		StartDate: c.StartDate,
		EndDate:   c.EndDate,
		Type:      c.Type,
	}
}

func dtoToPerson(dto Person) models.Person {
	return models.Person{
		ID:          dto.ID,
		Lastname:    dto.Lastname,
		Name:        dto.Name,
		Patronym:    dto.Patronym,
		DateOfBirth: dto.DateOfBirth,
		LocalRNOKPP: dto.LocalRNOKPP,
	}
}

func dtoToLegalPerson(dto LegalPerson) models.LegalPerson {
	return models.LegalPerson{
		ID:           dto.ID,
		Name:         dto.Name,
		ShortName:    dto.ShortName,
		RegisteredAt: dto.RegisteredAt,
		LocalKWED:    dto.LocalKWED,
		LocalEDRPOU:  dto.LocalEDRPOU,
	}
}

func dtoToContract(dto Contract) models.Contract {
	return models.Contract{
		ID:        dto.ID,
		Number:    dto.Number,
		StartDate: dto.StartDate,
		EndDate:   dto.EndDate,
		Type:      dto.Type,
	}
}
