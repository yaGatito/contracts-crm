package main

import (
	"context"
	sqlitedb "contract-service/adapters/db"
	"contract-service/models"
	"log"
	"time"
)

func main() {
	db, err := sqlitedb.NewSQLite("app.db")
	if err != nil {
		log.Fatal(err)
	}

	personRepo := sqlitedb.NewPersonRepository(db)
	// legalPersonRepo := sqlitedb.NewLegalPersonRepository(db)
	// contractRepo := sqlitedb.NewContractRepository(db)

	p := &models.Person{
		Lastname:    "Bomj",
		Name:        "Bomjara",
		Patronym:    "Bomjarovi4",
		DateOfBirth: time.Now(),
		LocalRNOKPP: 1231211231,
	}

	err = personRepo.Save(context.Background(), p)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(p.ID)

}
