package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllListBooks(c *fiber.Ctx) error {
	return c.JSON(ListBook)
}

func GetBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	for _, book := range ListBook {
		if book.ID == id {
			return c.JSON(book)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}

func CreateBook(c *fiber.Ctx) error {
	book := new(Book_Discription)

	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	book.ID = len(ListBook) + 1
	ListBook = append(ListBook, *book)

	return c.JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	bookUpdate := new(Book_Discription)
	if err := c.BodyParser(bookUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, Book := range ListBook {
		if Book.ID == id {
			Book.Title_Book = bookUpdate.Title_Book
			Book.Author = bookUpdate.Author
			ListBook[i] = Book
			return c.JSON(Book)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}

func DeleteBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	for i, Book := range ListBook {
		if Book.ID == id {
			ListBook = append(ListBook[:i], ListBook[i+1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}
