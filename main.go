package main

import (
	"krisboy-web-api/bootstrap"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	driver, err := bootstrap.GetAppDriver()
	if err != nil {
		panic(err.Error())
	}

	db, err := bootstrap.InitDatabase(driver.DB)

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": "hello world",
		})
	})

	app.Listen(":8080")
}
