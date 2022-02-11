package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Request struct {
	ID   *int    `json:"id" form:"id" query:"id"`
	Name *string `json:"name" form:"name" query:"name"`
}

type Response struct {
	ID   *int    `json:"id"`
	Name *string `json:"name"`
}

func echo(c *fiber.Ctx) error {

	query := c.Request().URI().QueryString()
	body := c.Body()
	// Check any input exists
	if len(query) == 0 && len(body) == 0 {
		return c.JSON(&Response{})
	}

	req := &Request{}

	// Query params have more priority in this case
	if len(body) > 0 {
		if err := c.BodyParser(req); err != nil {
			c.Status(400)
			return fiber.ErrBadRequest
		}
	}
	if err := c.QueryParser(req); err != nil {
		c.Status(400)
		return fiber.ErrBadRequest
	}
	response := &Response{ID: req.ID, Name: req.Name}
	return c.JSON(response)
}

func main() {
	app := fiber.New()

	app.Get("/echo", echo)
	app.Post("/echo", echo)

	log.Fatal(app.Listen(":3000"))
}
