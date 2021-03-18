package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/hosting-engine/envoy-proxy-manager-api/pkg/config"
	"log"
	"net/http"
)

func main() {
	settings := config.NewSettings("")
	app := fiber.New()

	app.Use(cors.New())
	app.Use(requestid.New())
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format:     "${time} - ${status} - ${method} ${path} ${latency}\n",
		TimeFormat: "2006-01-02T15:04:05-0700",
		TimeZone:   "Europe/Berlin",
	}))

	if settings.EnableAdminInterface {
		app.Use(filesystem.New(filesystem.Config{
			Root:         http.Dir(settings.AdminInterfacePath),
			Browse:       false,
			Index:        "index.html",
			NotFoundFile: "404.html",
			MaxAge:       3600,
		}))
	}


	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	if err := app.Listen(fmt.Sprintf(":%s", settings.Port)); err != nil {
		log.Fatal(fmt.Errorf("%w", err).Error())
	}
}
