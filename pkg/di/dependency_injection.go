package di

import (
	"github.com/anjush-bhargavan/golib_books/config"
	"github.com/anjush-bhargavan/golib_books/pkg/db"
	"github.com/anjush-bhargavan/golib_books/pkg/handler"
	"github.com/anjush-bhargavan/golib_books/pkg/repo"
	"github.com/anjush-bhargavan/golib_books/pkg/server"
	"github.com/anjush-bhargavan/golib_books/pkg/services"
)

// Init function load the configurations and initialize all instances
func Init() {
	config.LoadConfig()
	db := db.ConnectDB()
	bookRepo := repo.NewBookRepo(db)
	bookSvc := services.NewBookService(bookRepo)
	bookHandle := handler.NewBooksHandler(bookSvc)
	server.NewGrpcServer(bookHandle)

}
