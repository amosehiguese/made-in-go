package types

import (
	"github.com/amosehiguese/bookstore/storage"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name 			string 			`json:"name"`
	Author			string			`json:"author"`
	Publication		string			`json:"publication"`
}

func (b *Book) Create() *Book {
	db := storage.GetDB()
	db.Create(b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db := storage.GetDB()
	db.Find(&books)

	return books
}

func GetBookById(id uint) *Book {
	var book Book
	db := storage.GetDB()

	db.Where("ID=?", id).Find(&book)
	return &book
}

func (b *Book) Update(id uint) *Book {
	db := storage.GetDB()
	db.Where("id", id).Updates(b)
	return b
}

func (b *Book) Delete() {
	db := storage.GetDB()
	db.Delete(b)
}

