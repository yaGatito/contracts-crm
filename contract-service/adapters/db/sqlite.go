package sqlitedb

import (
	"contract-service/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLite() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.Person{},
		&models.LegalPerson{},
		&models.Contract{},
		&models.ContractPerson{},
		&models.ContractLegalPerson{},
	)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
