package controllers

import (
	"errors"
	"go-rest-book-api/mocks"
	"go-rest-book-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getBookById(id string) (*models.Book, *int, error) {
	for i, book := range mocks.MockedBooks {
		if book.ID == id {
			return &mocks.MockedBooks[i], &i, nil
		}
	}

	return nil, nil, errors.New("book not found")
}

// @Summary Get list of books
// @Description Get list of books
// @ID get-books
// @Produce  json
// @Success 200 {array} models.Book
// @Failure 500 {object} models.ErrorResponse
// @Router /books [get]
func GetBooks(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, mocks.MockedBooks)
}

// @Summary Get a book by ID
// @Description Get a book by ID
// @ID get-book-by-id
// @Produce  json
// @Param id path string true "Book ID"
// @Success 200 {object} models.Book
// @Failure 404 {object} models.ErrorResponse
// @Router /books/{id} [get]
func GetBook(ctx *gin.Context) {
	id := ctx.Param("id")

	book, _, err := getBookById(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Book not found", "status": http.StatusNotFound})
		return
	}

	ctx.IndentedJSON(http.StatusFound, book)
}

// @Summary Add a new book
// @Description Add a new book
// @ID add-book
// @Accept  json
// @Produce  json
// @Param book body models.Book true "Book object"
// @Success 201 {object} models.Book
// @Failure 400 {object} models.ErrorResponse
// @Router /books [post]
func AddBook(ctx *gin.Context) {
	var newBook models.Book

	if err := ctx.BindJSON(&newBook); err != nil {
		return
	}

	mocks.MockedBooks = append(mocks.MockedBooks, newBook)

	ctx.IndentedJSON(http.StatusCreated, newBook)
}

// @Summary Update book availability
// @Description Update book availability
// @ID update-book-availability
// @Produce  json
// @Param id path string true "Book ID"
// @Success 200 {object} models.Book
// @Failure 404 {object} models.ErrorResponse
// @Router /books/{id} [patch]
func UpdateBookAvailability(ctx *gin.Context) {
	id := ctx.Param("id")

	book, _, err := getBookById(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Book not found", "status": http.StatusNotFound})
		return
	}

	book.InStore = !book.InStore
	ctx.IndentedJSON(http.StatusOK, book)
}

// @Summary Remove a book
// @Description Remove a book
// @ID remove-book
// @Produce  json
// @Param id path string true "Book ID"
// @Success 200 {string} string
// @Failure 404 {object} models.ErrorResponse
// @Router /books/{id} [delete]
func RemoveBook(ctx *gin.Context) {
	id := ctx.Param("id")

	_, index, err := getBookById(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Book not found", "status": http.StatusNotFound})
		return
	}

	mocks.MockedBooks = append(mocks.MockedBooks[:*index], mocks.MockedBooks[*index+1:]...)
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Book removed successfully"})
}
