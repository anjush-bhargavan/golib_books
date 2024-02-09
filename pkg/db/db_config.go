package db

import (
	"log"
	"os"

	"github.com/anjush-bhargavan/golib_books/pkg/dom"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn, ok := os.LookupEnv("DB_Config")
	if !ok {
		log.Fatal("error loading database env")
	}

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// DB,err := sql.Open("postgres",dsn)
	if err != nil {
		log.Fatalf("error conncting to databse: %v", err.Error())
	}
	DB.AutoMigrate(&dom.Book{},
		&dom.Orders{})

	return DB
}
