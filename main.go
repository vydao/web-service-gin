package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID     int32   `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

var books = []book{
	{ID: 1, Title: "Blue Train", Author: "John Coltrane", Price: 56.99},
	{ID: 2, Title: "Jeru", Author: "Gerry Mulligan", Price: 17.99},
	{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Author: "Sarah Vaughan", Price: 39.99},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func addBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		log.Println("could not bind the received JSON")
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", addBook)

	router.Run("localhost:8080")
}
