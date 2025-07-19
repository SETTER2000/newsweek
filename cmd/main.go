package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	slogfiber "github.com/samber/slog-fiber"
	"github.com/setter2008/newsweek/config"
	"github.com/setter2008/newsweek/internal/pages"
	"github.com/setter2008/newsweek/pkg/logger"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()
	customLogger := logger.NewLogger(logConfig)
	app := fiber.New()

	app.Use(slogfiber.New(customLogger))
	app.Use(recover.New())

	pages.NewHandler(app)

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
