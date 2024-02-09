package repo

import (
	"github.com/anjush-bhargavan/golib_books/pkg/dom"
	"gorm.io/gorm"
)

// BookRepository for connecting db to BookRepoInter methods
type BookRepository struct {
	DB *gorm.DB
}

// NewBookRepo creates an instance of Book repo
func NewBookRepo(db *gorm.DB) *BookRepository {
	return &BookRepository{DB: db}
}

// CreateBook creates a newBook in database else returns error
func (u *BookRepository) CreateBook(Book *dom.Book) error {
	if err := u.DB.Create(Book).Error; err != nil {
		return err
	}
	return nil
}

// FindBookByID finds the Book in the database using ID
func (u *BookRepository) FindBookByID(id uint) (*dom.Book, error) {
	var Book *dom.Book
	err := u.DB.First(&Book, id).Error
	return Book, err
}

// UpdateBook updates Book details in the databse
func (u *BookRepository) UpdateBook(Book *dom.Book) error {
	err := u.DB.Save(&Book).Error
	return err
}

// DeleteBook deletes a Book using the id
func (u *BookRepository) DeleteBook(id uint) error {
	err := u.DB.Delete(&dom.Book{}, id).Error
	return err
}

// FindBookByEmail finds the Book by email
func (u *BookRepository) FindBookByName(name string) (*dom.Book, error) {
	var Book *dom.Book
	err := u.DB.Where("book_name = ?", name).First(&Book).Error
	return Book, err
}

// FindAllBooks finds the all the Book details in the database
func (u *BookRepository) FindAllBooks() ([]*dom.Book, error) {
	var Books []*dom.Book
	err := u.DB.Find(&Books).Error
	return Books, err
}
