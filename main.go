package main

import (
	"errors"
	"net/http"

	_ "go-rest-book-api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Book struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Published   int    `json:"published"`
	InStore     bool   `json:"in_store"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

var books = []Book{
	{
		ID:          "1",
		Title:       "The Great Gatsby",
		Description: "The story of the mysteriously wealthy Jay Gatsby and his love for the beautiful Daisy Buchanan.",
		Author:      "F. Scott Fitzgerald",
		Published:   1925,
		InStore:     false,
	},
	{
		ID:          "2",
		Title:       "To Kill a Mockingbird",
		Description: "The story of a young girl in a Southern town and her father, a lawyer who defends a black man accused of raping a white woman.",
		Author:      "Harper Lee",
		Published:   1960,
		InStore:     true,
	},
	{
		ID:          "3",
		Title:       "1984",
		Description: "A dystopian novel by George Orwell, and one of the best-known works of dystopian fiction.",
		Author:      "George Orwell",
		Published:   1949,
		InStore:     false,
	},
	{
		ID:          "4",
		Title:       "Moby-Dick",
		Description: "The story of the voyage of the whaling ship Pequod, commanded by Captain Ahab, who leads his crew on a hunt for the great whale Moby Dick.",
		Author:      "Herman Melville",
		Published:   1851,
		InStore:     true,
	},
	{
		ID:          "5",
		Title:       "Pride and Prejudice",
		Description: "The story follows the main character, Elizabeth Bennet, as she deals with issues of manners, upbringing, morality, education, and marriage in the society of the landed gentry of early 19th-century England.",
		Author:      "Jane Austen",
		Published:   1813,
		InStore:     true,
	},
	{
		ID:          "6",
		Title:       "The Catcher in the Rye",
		Description: "The story of Holden Caulfield's experiences in New York City in the days following his expulsion from Pencey Prep, a fictional college preparatory school.",
		Author:      "J.D. Salinger",
		Published:   1951,
		InStore:     true,
	},
	{
		ID:          "7",
		Title:       "Harry Potter and the Philosopher's Stone",
		Description: "The story of a young wizard, Harry Potter, and his friends Hermione Granger and Ron Weasley, who are students at Hogwarts School of Witchcraft and Wizardry.",
		Author:      "J.K. Rowling",
		Published:   1997,
		InStore:     true,
	},
	{
		ID:          "8",
		Title:       "The Lord of the Rings",
		Description: "The story begins in the Shire, where the hobbit Frodo Baggins inherits the Ring from Bilbo Baggins, his cousin and guardian.",
		Author:      "J.R.R. Tolkien",
		Published:   1954,
		InStore:     true,
	},
	{
		ID:          "9",
		Title:       "The Hobbit",
		Description: "The story of Bilbo Baggins, a hobbit who lives in Hobbiton, who is visited by the wizard Gandalf and 13 dwarves.",
		Author:      "J.R.R. Tolkien",
		Published:   1937,
		InStore:     true,
	},
	{
		ID:          "10",
		Title:       "The Adventures of Huckleberry Finn",
		Description: "The story of Huck Finn's adventures with his friend Tom Sawyer, and their journey down the Mississippi River on a raft.",
		Author:      "Mark Twain",
		Published:   1884,
		InStore:     true,
	},
}

// @Summary Get list of books
// @Description Get list of books
// @ID get-books
// @Produce  json
// @Success 200 {array} Book
// @Failure 500 {object} ErrorResponse
// @Router /books [get]
func getBooks(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, books)
}

// @Summary Add a new book
// @Description Add a new book
// @ID add-book
// @Accept  json
// @Produce  json
// @Param book body Book true "Book object"
// @Success 201 {object} Book
// @Failure 400 {object} ErrorResponse
// @Router /books [post]
func addBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)

	ctx.IndentedJSON(http.StatusCreated, newBook)
}

func getBookById(id string) (*Book, *int, error) {
	for i, book := range books {
		if book.ID == id {
			return &books[i], &i, nil
		}
	}

	return nil, nil, errors.New("Book not found")
}

// @Summary Get a book by ID
// @Description Get a book by ID
// @ID get-book-by-id
// @Produce  json
// @Param id path string true "Book ID"
// @Success 200 {object} Book
// @Failure 404 {object} ErrorResponse
// @Router /books/{id} [get]
func getBook(ctx *gin.Context) {
	id := ctx.Param("id")

	book, _, err := getBookById(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Book not found", "status": http.StatusNotFound})
		return
	}

	ctx.IndentedJSON(http.StatusFound, book)
}

// @Summary Update book availability
// @Description Update book availability
// @ID update-book-availability
// @Produce  json
// @Param id path string true "Book ID"
// @Success 200 {object} Book
// @Failure 404 {object} ErrorResponse
// @Router /books/{id} [patch]
func updateBookAvailability(ctx *gin.Context) {
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
// @Failure 404 {object} ErrorResponse
// @Router /books/{id} [delete]
func removeBook(ctx *gin.Context) {
	id := ctx.Param("id")

	_, index, err := getBookById(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Book not found", "status": http.StatusNotFound})
		return
	}

	books = append(books[:*index], books[*index+1:]...)
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Book removed successfully"})
}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:9090
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")

	{
		books := v1.Group("/books")
		{
			books.GET("", getBooks)
			books.POST("", addBook)
			books.GET(":id", getBook)
			books.PATCH(":id", updateBookAvailability)
			books.DELETE(":id", removeBook)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":9090")
}
