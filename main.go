package main

import (
	a "mongoapi/api"

	"github.com/gin-gonic/gin"
)

// estoy siguiendo el turo, aqui va la declaración de structs
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// Rellenar para poblar vars
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 500},
}

// definir bien puerto
func main() {
	router := gin.Default()
	router.GET("/albums", a.GetAlbums)
	router.POST("/albums", a.PostAlbums)
	router.GET("/albums/:id", a.GetAlbumByID)

	router.Run("127.0.0.1:8080")
}
