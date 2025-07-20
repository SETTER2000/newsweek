package pages

import "github.com/gofiber/fiber/v2"

type Handler struct {
	router fiber.Router
}

type Items struct {
	ID   int
	Name string
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

	names := []string{"Еда", "Животные", "Машины", "Спорт", "Музыка", "Технологии", "Прочее"}
	items := []Items{
		{ID: 1, Name: "Anton"},
		{ID: 2, Name: "Vasia"},
	}

	//names := []string{"Антон", "Вася"}
	data := struct {
		Names []string
		Items []Items
	}{Names: names, Items: items}
	return c.Render("page", data)
}
func (h *Handler) error(c *fiber.Ctx) error {
	//return fmt.Errorf("Error")
	return c.SendString("Error")
}
