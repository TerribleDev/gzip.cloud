package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

func main() {
	app := fiber.New()
	app.Use(compress.New())

	app.Use(compress.New(compress.Config{
		Level: compress.LevelDefault,
	}))
	// GET /api/register
	app.Get("/*", func(c *fiber.Ctx) error {
		println(c.Get("*"))
		resp, err := http.Get("https://" + c.Params("*"))
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		c.Response().Header.Add("Content-Type", resp.Header.Get("Content-Type"))
		c.Response().SetBodyStream(resp.Body, int(resp.ContentLength))
		return nil
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}
