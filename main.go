package main

import (
	"fiber/internal/delivery"
	"fiber/internal/repository"
	"fiber/internal/usescase"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	r := repository.NewUserRepository()
	u := usescase.NewUserUsecase(r)

	delivery.NewUserHandler(app, u)

	app.Listen(":3000")
}
