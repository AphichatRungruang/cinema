package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func getAllListMovies(c *fiber.Ctx) error {
	return c.JSON(ListMovies)
}

func getMovie(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	for _, movie := range ListMovies {
		if movie.ID == id {
			return c.JSON(movie)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}

func createMovie(c *fiber.Ctx) error {
	movie := new(Movie_Discription)

	if err := c.BodyParser(movie); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	movie.ID = len(ListMovies) + 1
	ListMovies = append(ListMovies, *movie)

	return c.JSON(movie)
}

func updateMovie(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	movieUpdate := new(Movie_Discription)
	if err := c.BodyParser(movieUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, movie := range ListMovies {
		if movie.ID == id {
			movie.Title_Movie = movieUpdate.Title_Movie
			movie.Rating = movieUpdate.Rating
			ListMovies[i] = movie
			return c.JSON(movie)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}

func deleteMovie(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	for i, movie := range ListMovies {
		if movie.ID == id {
			ListMovies = append(ListMovies[:i], ListMovies[i+1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}
