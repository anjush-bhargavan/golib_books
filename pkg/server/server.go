package server

import (
	"log"
	"net"
	pb "github.com/anjush-bhargavan/golib_books/pkg/pb"
	"github.com/anjush-bhargavan/golib_books/pkg/handler"
	"google.golang.org/grpc"
)

func NewGrpcServer(handlr *handler.BooksHandler) {
	log.Println("connecting to gRPC server")
	lis, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatal("error creating listener on port 8083")
	}

	grp := grpc.NewServer()
	pb.RegisterBookServicesServer(grp, handlr)

	log.Printf("listening on gRPC server 8083")
	if err := grp.Serve(lis); err != nil {
		log.Fatal("error connecting to gRPC server")

	}
}
