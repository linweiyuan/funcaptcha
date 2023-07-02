package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/linweiyuan/funcaptcha"
	"log"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		token, err := funcaptcha.GetOpenAIToken()
		if err != nil {
			return c.JSON(map[string]any{"error": err.Error()})
		}
		return c.JSON(token)
	})

	err := app.Listen(":3610")
	log.Printf("server running! port ===> 3610")

	if err != nil {
		log.Printf("启动失败! err info ===>", err)
	}

}
