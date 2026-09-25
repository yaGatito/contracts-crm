package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"path/filepath"

	sqlitedb "contract-service/adapters/db"
	httpadapter "contract-service/adapters/http"
	"contract-service/service"

	"github.com/gin-gonic/gin"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

type App struct {
	server *gin.Engine
}

func NewApp() *App {
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

	return &App{server: router}
}

func (a *App) Start(ctx context.Context) {
	go func() {
		if err := a.server.Run("127.0.0.1:8080"); err != nil {
			log.Printf("backend server stopped: %v", err)
		}
	}()
}

func main() {
	app := NewApp()

	if err := wails.Run(&options.App{
		Title:  "Contracts CRM",
		Width:  1500,
		Height: 1000,
		OnStartup: func(ctx context.Context) {
			app.Start(ctx)
		},
		AssetServer: &assetserver.Options{
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				frontendDir := filepath.Join(".", "dist")
				if _, err := os.Stat(filepath.Join(frontendDir, "index.html")); err == nil {
					http.FileServer(http.Dir(frontendDir)).ServeHTTP(w, r)
					return
				}
				w.Header().Set("Content-Type", "text/html; charset=utf-8")
				_, _ = w.Write([]byte("<html><body><h1>Contracts CRM Desktop</h1><p>Frontend bundle not found.</p></body></html>"))
			}),
		},
	}); err != nil {
		log.Fatal(err)
	}
}
