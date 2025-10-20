package delivery

import (
	"fiber/internal/entity"
	"fiber/internal/usescase"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	usescase usescase.UserUsecase
}

func NewUserHandler(app *fiber.App, u usescase.UserUsecase) {
	handler := &userHandler{usescase: u}

	app.Get("/users", handler.GetUsers)
	app.Post("/users", handler.CreateUsers)
}

func (h *userHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.usescase.GetUsers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(users)
}

func (h *userHandler) CreateUsers(c *fiber.Ctx) error {
	var user entity.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid body"})
	}

	if err := h.usescase.CreateUsers(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"message": "User Created"})
}
