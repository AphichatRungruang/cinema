package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type Book_Discription struct {
	ID         int    `json:"id"`
	Title_Book string `json:"book"`
	Author     string `json:"author"`
}

var ListBook []Book_Discription

func main() {

	template_engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: template_engine,
	})

	ListBook = append(ListBook, Book_Discription{ID: 1, Title_Book: "Golang", Author: "Me"})
	ListBook = append(ListBook, Book_Discription{ID: 2, Title_Book: "API", Author: "Me"})

	app.Get("/Book", GetAllListBooks)
	app.Get("/Book/:id", GetBook)
	app.Post("/Book", CreateBook)
	app.Put("/Book/:id", UpdateBook)
	app.Delete("/Book/:id", DeleteBook)

	app.Post("/upload", uploadFile)
	app.Get("/test_html", testHTML)

	app.Listen(":8080")
}

func uploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("image") //name Request

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	err = c.SaveFile(file, "./uploads/"+file.Filename) // nameFoder

	if (err) != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendString("File upload complet!")
}
func testHTML(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Tital": "Hello,World!",
		"Name":  "FLUKEFY",
	})
}
