package handler

import (
	"context"

	pb "github.com/anjush-bhargavan/golib_books/pkg/pb"
	services_interfaces "github.com/anjush-bhargavan/golib_books/pkg/services/interfaces"
)

type BooksHandler struct {
	pb.BookServicesServer
	svc services_interfaces.BookServiceInter
}

func NewBooksHandler(svc services_interfaces.BookServiceInter) *BooksHandler {
	return &BooksHandler{
		svc: svc,
	}
}

func (b *BooksHandler) CreateBook(ctx context.Context, p *pb.BookModel) (*pb.BookResponse, error) {
	result, err := b.svc.CreateBookService(p)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (b *BooksHandler) FetchBookByID(ctx context.Context, p *pb.BookID) (*pb.BookModel, error) {
	result, err := b.svc.FindBookByIDService(p)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (b *BooksHandler) FetchBookByName(ctx context.Context, p *pb.BookName) (*pb.BookModel, error) {
	result, err := b.svc.FindBookByNameService(p)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (b *BooksHandler) FetchAllBooks(ctx context.Context, p *pb.NoParam) (*pb.BookList, error) {
	result, err := b.svc.FindAllBooksService(p)
	if err != nil {
		return nil, err
	}

	return result, nil
}
