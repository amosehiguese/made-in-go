package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/amosehiguese/bookstore/types"
	"github.com/go-chi/chi/v5"
)

func parseBody(r *http.Request, st any) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, st)
}


func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book types.Book
	parseBody(r, &book)

	newBook := book.Create()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(newBook)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := types.GetAllBooks()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(books)
}

func RetrieveBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "bookId"))
	if err != nil{
		log.Fatal(err)
	}

	book := types.GetBookById(uint(id))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(book)
}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "bookId"))
	if err != nil {
		log.Fatal(err)
	}

	var book types.Book
	parseBody(r, &book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	book.Update(uint(id))
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "bookId"))
	if err != nil {
		log.Fatal(err)
	}

	book := types.GetBookById(uint(id))
	book.Delete()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}
