package main

import (
	"github.com/gofiber/fiber/v2"
)

type Movie_Discription struct {
	ID          int    `json:"id"`
	Title_Movie string `json:"movie"`
	Rating      int    `json:"author"`
}

var ListMovies []Movie_Discription

func main() {
	app := fiber.New()

	ListMovies = append(ListMovies, Movie_Discription{ID: 1, Title_Movie: "Avenger EndGame", Rating: 10.0})
	ListMovies = append(ListMovies, Movie_Discription{ID: 2, Title_Movie: "Iron Man", Rating: 9.0})

	app.Get("/Movie", getAllListMovies)
	app.Get("/Movie/:id", getMovie)
	app.Post("/Movie", createMovie)
	app.Put("/Movie/:id", updateMovie)
	app.Delete("/Movie/:id", deleteMovie)

	app.Listen(":8080")
}
