package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"pustaka-api/book"
	"pustaka-api/helper"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (handler bookHandler) PostBookHandler(c *gin.Context) {
	//get body
	var bookRequest book.BookRequest
	err := c.ShouldBindJSON(&bookRequest)

	var jsonErr *json.UnmarshalTypeError
	if errors.As(err, &jsonErr) {
		// log.Println("Json binding error", jsonErr)
		ermsg := fmt.Sprintf("error Field %s value %s must be %s ", jsonErr.Field, jsonErr.Value, jsonErr.Type)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": ermsg,
		})
		return
	}
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	book, err := handler.bookService.Create(bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	//response
	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func (handler bookHandler) GetBookHandler(c *gin.Context) {

	books, err := handler.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	var booksResp []book.BookResponse

	for _, b := range books {
		bookResp := book.BookResponse{
			ID:          b.ID,
			Title:       b.Title,
			Description: b.Description,
			Rating:      b.Rating,
			Price:       b.Price,
		}

		booksResp = append(booksResp, bookResp)
	}

	res := helper.BuildResponse(true, "OK!", books)
	c.JSON(http.StatusOK, res)
}

func (handler bookHandler) GeBookID(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	book, err := handler.bookService.FindByID(id)
	if err != nil {
		res := helper.BuildErrorResponse("Failed Get Params", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})

}

func (handler bookHandler) PutBookHandler(c *gin.Context) {
	var bookRequest book.BookFindByID
	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

		}
		c.JSON(http.StatusOK, gin.H{
			"data": errorMessages,
		})
		return
	}

	book, err := handler.bookService.FindByID(bookRequest.ID)

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}
