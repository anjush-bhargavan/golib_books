package services

import (
	"github.com/anjush-bhargavan/golib_books/pkg/dom"
	pb "github.com/anjush-bhargavan/golib_books/pkg/pb"
	repo_interfaces "github.com/anjush-bhargavan/golib_books/pkg/repo/interfaces"
	services_interfaces "github.com/anjush-bhargavan/golib_books/pkg/services/interfaces"
)

type BookService struct {
	repo repo_interfaces.BookRepoInter
}

func NewBookService(repo repo_interfaces.BookRepoInter) services_interfaces.BookServiceInter {
	return &BookService{
		repo: repo,
	}
}

func (b *BookService) CreateBookService(p *pb.BookModel) (*pb.BookResponse, error) {
	book := &dom.Book{
		BookName:     p.BookName,
		Author:       p.Author,
		NoOfCopies:   uint64(p.NoOfCopies),
		Year:         p.Year,
		Publications: p.Publications,
		Category:     p.Category,
		Description:  p.Description,
	}

	err := b.repo.CreateBook(book)
	if err != nil {
		return nil, err
	}
	return &pb.BookResponse{
		Status:  "Success",
		Error:   "",
		Message: "book added successfully",
	}, nil
}

func (b *BookService) FindBookByIDService(p *pb.BookID) (*pb.BookModel, error) {
	book, err := b.repo.FindBookByID(uint(p.Id))
	if err != nil {
		return nil, err
	}
	return &pb.BookModel{
		BookName: book.BookName,
		Author: book.Author,
		NoOfCopies: uint32(book.NoOfCopies),
		Year: book.Year,
		Publications: book.Publications,
		Category: book.Category,
		Description: book.Description,
	},nil
}

func (b *BookService) FindBookByNameService(p *pb.BookName) (*pb.BookModel, error) {
	book, err := b.repo.FindBookByName(p.Name)
	if err != nil {
		return nil, err
	}
	return &pb.BookModel{
		BookName: book.BookName,
		Author: book.Author,
		NoOfCopies: uint32(book.NoOfCopies),
		Year: book.Year,
		Publications: book.Publications,
		Category: book.Category,
		Description: book.Description,
	},nil
}

func (b *BookService) FindAllBooksService(p *pb.NoParam) (*pb.BookList,error) {
	result,err := b.repo.FindAllBooks()
	if err != nil {
		return nil,err
	}
	var books []*pb.BookModel
	for _,book := range result {
		books = append(books,&pb.BookModel{
			BookName: book.BookName,
			Author: book.Author,
			NoOfCopies: uint32(book.NoOfCopies),
			Year: book.Year,
			Publications: book.Publications,
			Category: book.Category,
			Description: book.Description,
		})
	}

	return &pb.BookList{
		Books: books,
	},nil
}
