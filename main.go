package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/pustaka?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database error connection")
	}
	/**
	db.AutoMigrate(&book.Book{})

	------------------- CRUD ------------------

	//Create
	book := book.Book{}
	book.Title = "Pemrograman Golang"
	book.Description = "Sevuah buku untuk belajar C#"
	book.Price = 1120000
	book.Rating = 4

	err = db.Create(&book).Error
	if err != nil {
		panic("Error creating book data")
	}

	//GET All Book
	var books []book.Book
	err = db.Find(&books).Error
	if err != nil {
		panic("Error get all book")
	}
	for i, value := range books {
		fmt.Println("Buku", i+1, "Judul :", value.Title)
	}

	//GET Single Book
	var book book.Book
	err = db.Find(&book, 1).Error
	if err != nil {
		panic("Error get single book")
	}
	fmt.Println(book.Title)


	//GET All Book Using Condition
	var books []book.Book
	value := "catatan"
	err = db.Where("title LIKE ?", "%"+value+"%").Find(&books).Error
	if err != nil {
		panic("Error get all book")
	}
	for i, value := range books {
		fmt.Println("Buku", i+1, "Judul :", value.Title)
	}


	//Update For Single Book
	var book = book.Book{ID: 31}
	rows := db.First(&book).RowsAffected
	if rows == 0 {
		panic("Book not found, try again")
	}
	book.Title = "Catatan si Natan"
	err = db.Save(&book).Error
	if err != nil {
		panic("Update book failed")
	}
	fmt.Println(book)

	//Delete Book
	var book = book.Book{ID: 3}
	rows := db.First(&book).RowsAffected
	if rows == 0 {
		panic("Book not found, try again")
	}
	err = db.Delete(&book).Error
	if err != nil {
		panic("Delete book failed")
	}
	fmt.Println("Berhasil hapus buku!")
	*/

	// router := gin.Default()
	// v1 := router.Group("v1")

	// v1.GET("/", handlers.RootHandler)
	// v1.GET("/hello", handlers.HelloHandler)
	// v1.GET("/book/:id", handlers.GetBookHandler)
	// v1.GET("/query", handlers.QueryHandler)
	// v1.POST("/book", handlers.CreateBookHandler)

	// router.Run()
}
