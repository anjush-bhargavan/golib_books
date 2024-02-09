package repo_interfaces

import "github.com/anjush-bhargavan/golib_books/pkg/dom"

// BookRepo defines the methods in BookRepo
type BookRepoInter interface {
	CreateBook(Book *dom.Book) error
	FindBookByID(id uint) (*dom.Book, error)
	UpdateBook(Book *dom.Book) error
	FindBookByName(name string) (*dom.Book, error)
	DeleteBook(id uint) error
	FindAllBooks() ([]*dom.Book, error)
}
