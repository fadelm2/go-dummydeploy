package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"

	"github.com/spf13/viper"
)

func loadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("⚠️ Config file tidak ditemukan, pakai default.")
	}
}

func main() {
	loadConfig()

	app := fiber.New()

	// health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "run jalan",
		})
	})
	// gateway
	app.Get("/gateway", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "gateway running",
		})
	})

	port := viper.GetInt("app.port")
	fmt.Printf("🚀 Server running on port %d\n", port)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
