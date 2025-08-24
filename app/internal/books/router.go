package books

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type BookRouter struct {
	app *fiber.App
	uc UseCaseInt
}

func NewBookRouter (app *fiber.App, usecase UseCaseInt) *BookRouter {
	return &BookRouter{
		app: app,
		uc: usecase,
	}
}

func (r *BookRouter) Register() {
	api := r.app.Group("/api") // /api/*
	books := api.Group("/books") // /api/books/*

	books.Get("/:id", r.GetBook)
	books.Post("/", r.PostBook)
	books.Put("/", r.PutBook)
	books.Delete("/:id", r.DeleteBook)
}

func (r *BookRouter) GetBook(c *fiber.Ctx) error {
	const op = "book.router.GetBook"

	idInt, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fmt.Errorf("%s: %w", op, ErrIncorrectParameter)
	}

	book, err := r.uc.GetBook(c.UserContext(), idInt)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	response := map[string]interface{}{
		"book": book,
		"name": "Ivan",
	}
	return c.JSON(response)
}

func (r *BookRouter) PostBook(c *fiber.Ctx) error {
	const op = "book.router.GetBook"

	var book BookDTO
    if err := c.BodyParser(&book); err != nil {
        return err
    }

	id, err := r.uc.CreateBook(c.UserContext(), book)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

    return c.JSON(fiber.Map{"id": id})
}

func (r *BookRouter) PutBook(c *fiber.Ctx) error {
	const op = "book.router.UpdateBook"

	var book BookResponse
    if err := c.BodyParser(&book); err != nil {
        return err
    }

	id, err := r.uc.UpdateBook(c.UserContext(), book)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}


    return c.JSON(fiber.Map{"id": id})
}

func (r *BookRouter) DeleteBook(c *fiber.Ctx) error {
	const op = "book.router.DeleteBook"

	idInt, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fmt.Errorf("%s: %w", op, ErrIncorrectParameter)
	}

	id, err := r.uc.RemoveBook(c.UserContext(), idInt)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return c.JSON(fiber.Map{"id": id})
}
