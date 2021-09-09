package main

import (
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

func getBookings(c *gin.Context) {
	c.IndentedJSON(http.StatusOK)
}

func main() {
}
