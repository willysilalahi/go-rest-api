package handlers

import (
	"go-rest-api/book"
	"net/http"

	"github.com/gin-gonic/gin"
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
	c.JSON(http.StatusOK, result)
}

func (h *bookHandler) GetSingleBookHandler(c *gin.Context) {
	result, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *bookHandler) CreateBookHandler(c *gin.Context) {
	var bookRequest book.BookRequest
	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {

	} else {
		c.JSON(http.StatusOK, "Aman bang")
	}
	// if err != nil {
	// 	var errorMessages []string
	// 	for _, e := range err.(validator.ValidationErrors) {
	// 		error := fmt.Sprintf("Error on field %s, condition : %s", e.Field(), e.ActualTag())
	// 		errorMessages = append(errorMessages, error)
	// 	}
	// 	c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
	// 	return
	// }

	// result, error := h.bookService.Create(bookRequest)

	// if error != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"errors": error})
	// 	return
	// }
	// c.JSON(http.StatusOK, result)
}
