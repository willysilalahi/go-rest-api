package handlers

import (
	"fmt"
	"go-rest-api/book"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Mantap nakku, bibik bangga"})
}

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Calvin Server",
		"desc": "A Server Cyber Security with tall body",
	})
}

func GetBookHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"name": "Hello Jhon",
		"desc": "Kamu berusaha ambil buku dengan id " + id,
	})
}

func QueryHandler(c *gin.Context) {
	id := c.Query("id")
	title := c.Query("title")
	c.JSON(http.StatusOK, gin.H{
		"name":  "Hello Jhon",
		"id":    id,
		"title": title,
	})
}

func CreateBookHandler(c *gin.Context) {
	var book book.BookInput
	err := c.ShouldBindJSON(&book)

	if err != nil {
		var errorMessages []string
		for _, e := range err.(validator.ValidationErrors) {
			error := fmt.Sprintf("Error on field %s, condition : %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, error)
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title":      book.Title,
		"price":      book.Price,
		"short_desc": book.ShortDesc,
	})
}
