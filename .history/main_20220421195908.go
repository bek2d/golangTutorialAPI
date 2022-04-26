package main

import (
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

func getBooks(c *gin.Context) {
	c.JSON(200, books)
}

func createBook(c *gin.Context) {
	var newBook Book
	c.BindJSON(&newBook)
	books = append(books, newBook)
	c.JSON(200, gin.H{"message": "Book created successfully!", "book": newBook})
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.Run(":8080")
}
