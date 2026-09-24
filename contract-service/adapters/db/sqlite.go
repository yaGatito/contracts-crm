package sqlitedb

import (
	"contract-service/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLite(path string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.Person{},
		&models.LegalPerson{},
		&models.Contract{},
	)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
