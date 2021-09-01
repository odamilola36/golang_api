package service

import (
	"fmt"
	"github.com/mashingan/smapping"
	"github.com/odamilola36/golang_api/dto"
	"github.com/odamilola36/golang_api/entity"
	"github.com/odamilola36/golang_api/repositories"
	"log"
)

type BookService interface {
	Insert(bookDto dto.BookCreateDTO) entity.Book
	Update(bookDto dto.BookUpdateDTO) entity.Book
	Delete(b entity.Book)
	All() []entity.Book
	FindById(bookId uint64) entity.Book
	IsAllowedToEdit(userId string, bookId uint64) bool
}

type bookService struct {
	bookRepository repositories.BookRepository
}

func NewBookService(bookRepository repositories.BookRepository) BookService {
	return &bookService{
		bookRepository: bookRepository,
	}
}

func (b bookService) Insert(bookDto dto.BookCreateDTO) entity.Book {
	book := entity.Book{}
	err := smapping.FillStruct(book, smapping.MapFields(bookDto))
	if err != nil {
		log.Fatal("Failed to map %v", err)
	}
	res := b.bookRepository.InsertBook(book)
	return res
}

func (b bookService) Update(bookDto dto.BookUpdateDTO) entity.Book {
	book := entity.Book{}
	err := smapping.FillStruct(book, smapping.MapFields(bookDto))
	if err != nil {
		log.Fatal("Failed to map %v", err)
	}
	res := b.bookRepository.InsertBook(book)
	return res
}

func (b bookService) Delete(bo entity.Book) {
	b.bookRepository.DeleteBook(bo)
}

func (b bookService) All() []entity.Book {
	return b.bookRepository.AllBook()
}

func (b bookService) FindById(bookId uint64) entity.Book {
	return b.bookRepository.FindBookById(bookId)
}

func (b bookService) IsAllowedToEdit(userId string, bookId uint64) bool {
	book := b.bookRepository.FindBookById(bookId)
	id := fmt.Sprintf("%v", book.UserId)
	return id == userId
}
