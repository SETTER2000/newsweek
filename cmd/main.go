package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/setter2008/newsweek/config"
	"github.com/setter2008/newsweek/internal/pages"
	"log"
)

func main() {
	config.Init()
	dbConf := config.NewDatabaseConfig()
	log.Println(dbConf)

	app := fiber.New()
	app.Use(recover.New())

	pages.NewHandler(app)

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
