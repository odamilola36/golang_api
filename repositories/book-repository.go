package repositories

import (
	"github.com/odamilola36/golang_api/entity"
	"gorm.io/gorm"
)

type BookRepository interface {
	InsertBook(b entity.Book) entity.Book
	UpdateBook(b entity.Book) entity.Book
	DeleteBook(b entity.Book)
	AllBook() []entity.Book
	FindBookById(bookId uint64) entity.Book
}

type bookConnection struct {
	bookConnection *gorm.DB
}

func NewBookRepository(dbConn *gorm.DB) BookRepository {
	return &bookConnection{
		bookConnection: dbConn,
	}
}

func (db bookConnection) InsertBook(b entity.Book) entity.Book {
	db.bookConnection.Save(&b)
	db.bookConnection.Preload("User").Find(&b)
	return b
}

func (db bookConnection) UpdateBook(b entity.Book) entity.Book {
	db.bookConnection.Save(&b)
	db.bookConnection.Preload("User").Find(&b)
	return b
}

func (db bookConnection) DeleteBook(b entity.Book) {
	db.bookConnection.Delete(&b)
}

func (db bookConnection) AllBook() []entity.Book {
	var books []entity.Book
	db.bookConnection.Preload("User").Find(&books)
	return books
}

func (db bookConnection) FindBookById(bookId uint64) entity.Book {
	var book entity.Book
	db.bookConnection.Preload("User").Find(&book, bookId)
	return book
}
