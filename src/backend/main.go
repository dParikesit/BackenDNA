package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/kristabdi/Tubes3_13520058/handlers"
	"github.com/kristabdi/Tubes3_13520058/utils"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env load failed")
	}

	app := fiber.New()
	utils.ConnectSupabase()

	api := app.Group("/api")
	api.Post("/insert", handlers.DiseaseInsert)
	api.Post("/match/:algo", handlers.DiseaseMatch)
	api.Post("/history", handlers.HistoryQuery)

	app.Static("/", "./frontend/build")

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	err = app.Listen(":" + port)
	if err != nil {
		log.Fatalln("Server start error")
	}
}
