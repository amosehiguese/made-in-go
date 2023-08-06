package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Movie struct {
	Id					string			`json:"id"`
	ISBN				string			`json:"isbn"`
	Title				string			`json:"title"`
	Director			*Director		`json:"director"`
}

type Director struct{
	FirstName			string 		`json:"firstname"`
	LastName			string		`json:"lastname"`
}

var movies []Movie

func init() {
	movies = []Movie{
		{
			Id: "1",
			ISBN: "1234-12340-1234",
			Title: "Game of Thrones",
			Director: &Director{
				FirstName: "luke",
				LastName: "mike",
			},
		},
		{
			Id: "2",
			ISBN: "1234-12340-1234",
			Title: "Waters",
			Director: &Director{
				FirstName: "luke",
				LastName: "mike",
			},
		},
		{
			Id: "3",
			ISBN: "1234-12340-1234",
			Title: "White House",
			Director: &Director{
				FirstName: "luke",
				LastName: "mike",
			},
		},
		{
			Id: "4",
			ISBN: "1234-12340-1234",
			Title: "Exams",
			Director: &Director{
				FirstName: "luke",
				LastName: "mike",
			},
		},
		{
			Id: "5",
			ISBN: "1234-12340-1234",
			Title: "Kings",
			Director: &Director{
				FirstName: "luke",
				LastName: "mike",
			},
		},
		{
			Id: "6",
			ISBN: "1234-12340-1234",
			Title: "Life Out",
			Director: &Director{
				FirstName: "luke",
				LastName: "mike",
			},
		},
	}
}

func main() {
	r := chi.NewRouter()

	r.Get("/movies", getMovies)
	r.Get("/movies/{id}", retrieveMovie)
	r.Post("/movies", createMovie)
	r.Patch("/movies/{id}", updateMovie)
	r.Delete("/movies/{id}", deleteMovies)

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))

}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func retrieveMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")

	for _, movie := range movies {
		if movie.Id == id {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)

	movie.Id = strconv.Itoa(rand.Intn(10000))
	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")

	var updateMovie Movie
	json.NewDecoder(r.Body).Decode(&updateMovie)

	for _, movie := range movies {
		if movie.Id == id {
			movie.ISBN = updateMovie.ISBN
			movie.Title = updateMovie.Title
			movie.Director = updateMovie.Director

			json.NewEncoder(w).Encode(movie)
			return
		}
	}

}

func deleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	for idx, movie := range movies {
		if movie.Id == id {
			movies = append(movies[:idx], movies[idx + 1:]... )
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
}
