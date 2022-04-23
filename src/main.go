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
		log.Println("Env load failed")
	}

	app := fiber.New()
	utils.ConnectSupabase()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	api := app.Group("/api")
	api.Post("/insert/:name/:sequence", handlers.DiseaseInsert)
	api.Post("/match/:algo", handlers.DiseaseMatch)
	api.Get("/history", handlers.HistoryQuery)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	err = app.Listen(":" + port)
	if err != nil {
		log.Fatalln("Server start error")
	}
}
