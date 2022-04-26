package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// "errors"
)

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []Book{
	{ID: "1", Title: "Book 1", Author: "Author 1", Quantity: 1},
	{ID: "2", Title: "Book 2", Author: "Author 2", Quantity: 2},
	{ID: "3", Title: "Book 3", Author: "Author 3", Quantity: 3},
}

func getBookByID(id string) (*Book, error) {
	for _, book := range books {
		if book.ID == id {
			return &book, nil
		}
	}
	return nil, nil
}

func bookByID(c *gin.Context) {
	id := c.Params.ByName("id")
	book, err := getBookByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(200, book)
}

func getBooks(c *gin.Context) {
	c.JSON(200, books)
}

func createBook(c *gin.Context) {
	var newBook Book
	c.BindJSON(&newBook)
	books = append(books, newBook)
	c.JSON(http.StatusCreated, books)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.GET("/books/:id", bookByID)
	router.Run(":8080")
}
