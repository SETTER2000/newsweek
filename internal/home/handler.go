package home

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/setter2008/newsweek/pkg/tadapter"
	"github.com/setter2008/newsweek/views"
)

type Handler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

type Items struct {
	ID   int
	Name string
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := &Handler{
		router:       router,
		customLogger: customLogger,
	}

	h.router.Get("/", h.home)
	h.router.Get("/error", h.error)
}

func (h *Handler) home(c *fiber.Ctx) error {
	//names := []string{"Еда", "Животные", "Машины", "Спорт", "Музыка", "Технологии", "Прочее"}
	//data := struct {
	//	Names []string
	//}{Names: names}
	component := views.Main()
	return tadapter.Render(c, component)
}
func (h *Handler) error(c *fiber.Ctx) error {
	h.customLogger.Info().
		Bool("isAdmin", true).
		Str("email", "a@a.ru").
		Int("id", 10).
		Msg("Инфо")

	return c.SendString("Error")
}
