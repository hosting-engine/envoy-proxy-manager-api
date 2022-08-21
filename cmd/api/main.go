package main

import (
	"fmt"
	"github.com/hosting-engine/envoy-proxy-manager-api/pkg/utils"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/hosting-engine/envoy-proxy-manager-api/pkg/config"
)

func main() {
	app := fiber.New(fiber.Config{
		ReadTimeout:           time.Duration(60) * time.Second,
		WriteTimeout:          time.Duration(60) * time.Second,
		DisableKeepalive:      true,
		DisableStartupMessage: true,
		CaseSensitive:         true,
		Concurrency:           0,
		AppName:               "Envoy Proxy Manager API",
	})

	app.Use(cors.New())
	app.Use(requestid.New())
	app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("Envoy Proxy Manager API v%s", utils.GetVersion()))
	})

	if err := app.Listen(fmt.Sprintf(":%s", config.S.Port)); err != nil {
		log.Fatal(fmt.Errorf("%w", err).Error())
	}
}
