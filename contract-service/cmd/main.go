package main

import (
	"log"

	sqlitedb "contract-service/adapters/db"
	httpadapter "contract-service/adapters/http"
	"contract-service/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := sqlitedb.NewSQLite()
	if err != nil {
		log.Fatal(err)
	}

	personRepo := sqlitedb.NewPersonRepository(db)
	legalPersonRepo := sqlitedb.NewLegalPersonRepository(db)
	contractRepo := sqlitedb.NewContractRepository(db)
	contractLegalPersonRepo := sqlitedb.NewContractLegalPersonRepository(db)
	contractPersonRepo := sqlitedb.NewContractPersonRepository(db)

	svc := service.NewContractService(
		personRepo,
		legalPersonRepo,
		contractRepo,
		contractLegalPersonRepo,
		contractPersonRepo,
	)

	router := gin.Default()
	httpadapter.NewHandlers(svc).Register(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
