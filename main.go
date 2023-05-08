package main

import (
	"go-rest-api/book"
	"go-rest-api/handlers"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/pustaka?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database error connection")
	}

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handlers.NewHandler(bookService)

	router := gin.Default()
	v1 := router.Group("v1")

	v1.GET("/book", bookHandler.GetAllBookHandler)
	v1.GET("/book/:id", bookHandler.GetSingleBookHandler)
	v1.POST("/book", bookHandler.CreateBookHandler)
	v1.PATCH("/book/:id", bookHandler.UpdateBookHandler)
	v1.DELETE("/book/:id", bookHandler.DeleteBookHandler)

	router.Run()
}
