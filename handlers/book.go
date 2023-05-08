package handlers

import (
	"fmt"
	"go-rest-api/book"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) GetAllBookHandler(c *gin.Context) {
	result, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	var booksResponse []book.BookResponse
	for _, b := range result {
		booksResponse = append(booksResponse, refactoring(b))
	}

	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (h *bookHandler) GetSingleBookHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := h.bookService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	bookReponse := refactoring(result)
	c.JSON(http.StatusOK, gin.H{
		"data": bookReponse,
	})
}

func (h *bookHandler) CreateBookHandler(c *gin.Context) {
	var bookRequest book.BookRequest
	err := c.ShouldBindJSON(&bookRequest)

	if err != nil {
		var errorMessages []string
		for _, e := range err.(validator.ValidationErrors) {
			error := fmt.Sprintf("Error on field %s, condition : %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, error)
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	}

	result, error := h.bookService.Create(bookRequest)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": error})
		return
	}
	bookResponse := refactoring(result)
	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (h *bookHandler) UpdateBookHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var bookRequest book.BookRequest
	err := c.ShouldBindJSON(&bookRequest)

	if err != nil {
		var errorMessages []string
		for _, e := range err.(validator.ValidationErrors) {
			error := fmt.Sprintf("Error on field %s, condition : %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, error)
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	}

	result, error := h.bookService.Update(id, bookRequest)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": error})
		return
	}
	bookResponse := refactoring(result)
	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (h *bookHandler) DeleteBookHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
}

func refactoring(b book.Book) book.BookResponse {
	bookResponse := book.BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Description: b.Description,
		Price:       b.Price,
		Rating:      b.Rating,
	}
	return bookResponse
}
