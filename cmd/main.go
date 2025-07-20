package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	slogfiber "github.com/samber/slog-fiber"
	"github.com/setter2008/newsweek/config"
	"github.com/setter2008/newsweek/internal/pages"
	"github.com/setter2008/newsweek/pkg/logger"
	"strings"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()
	customLogger := logger.NewLogger(logConfig)
	engine := html.New("./html", ".html")
	engine.AddFuncMap(map[string]interface{}{
		"ToUpper": func(c string) string { return strings.ToUpper(c) },
	})

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Отдача статических файлов из папки "public"
	app.Static("/public", "./public")

	app.Use(slogfiber.New(customLogger))
	app.Use(recover.New())

	pages.NewHandler(app)

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
