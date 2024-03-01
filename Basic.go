package main

import (
	"github.com/gofiber/fiber/v2"
)

type Book_Discription struct {
	ID         int    `json:"id"`
	Title_Book string `json:"book"`
	Author     string `json:"author"`
}

var ListBook []Book_Discription

func main() {
	app := fiber.New()

	ListBook = append(ListBook, Book_Discription{ID: 1, Title_Book: "Golang", Author: "Me"})
	ListBook = append(ListBook, Book_Discription{ID: 2, Title_Book: "API", Author: "Me"})

	app.Get("/Book", GetAllListBooks)
	app.Get("/Book/:id", GetBook)
	app.Post("/Book", CreateBook)
	app.Put("/Book/:id", UpdateBook)
	app.Delete("/Book/:id", DeleteBook)

	app.Listen(":8080")
}
