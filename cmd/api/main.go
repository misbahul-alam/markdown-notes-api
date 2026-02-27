package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/misbahul-alam/markdown-notes-api/internal/config"
	"github.com/misbahul-alam/markdown-notes-api/internal/model"
	"github.com/misbahul-alam/markdown-notes-api/internal/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	config.Load()
	config.ConnectDatabase()
	err = config.DB.AutoMigrate(
		&model.User{},
		&model.Note{},
	)
	if err != nil {
		log.Fatalf("Error auto migrating users: %v", err)
		return
	}
	r := gin.Default()
	routes.SetupRoutes(r)

	if err = r.Run(":" + config.C.AppPort); err != nil {
		log.Fatalf("Error starting server: %s", err)
		return
	}
}
