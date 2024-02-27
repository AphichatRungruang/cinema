package movie

import (
	"fmt"
)

func init() {

}
func ReviewMovie(name string, rating float64) {

	fmt.Printf("I reviewed %s and it's rating is %f\n", name, rating)
}
func FindMovieName(imdbID string) string {
	switch imdbID {

	case "tt4154796":
		return "Avenger EndGame"
	case "tt1825683":
		return "Black Panter"
	}
	return "not found."

}
