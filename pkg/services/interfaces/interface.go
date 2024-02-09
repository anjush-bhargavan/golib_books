package services_interfaces

import (
	pb "github.com/anjush-bhargavan/golib_books/pkg/pb"
)

type BookServiceInter interface {
	// FindAllBooksService(p *pb.FetchAll) ([]*entities.Book, error)
	CreateBookService(p *pb.BookModel) (*pb.BookResponse, error)
	// EditBookService(p *pb.BookCreate) (*pb.BookCreate,error)
	FindBookByNameService(p *pb.BookName) (*pb.BookModel,error)
	FindBookByIDService(p *pb.BookID) (*pb.BookModel,error)
	FindAllBooksService(p *pb.NoParam) (*pb.BookList,error)
	// DeleteBookService(p *pb.ID) (*pb.CommonResponse,error)
}
