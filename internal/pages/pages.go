package pages

import "github.com/gofiber/fiber/v2"

type Handler struct {
	router fiber.Router
}

func NewHandler(router fiber.Router) {
	h := &Handler{
		router: router,
	}
	api := router.Group("/api")
	api.Get("/", h.home)
	api.Get("/error", h.error)
}

func (h *Handler) home(c *fiber.Ctx) error {
	return c.SendString("Hello")
}
func (h *Handler) error(c *fiber.Ctx) error {
	return c.SendString("Error")
}
