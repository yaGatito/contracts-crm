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

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization, Accept, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	httpadapter.NewHandlers(svc).Register(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
