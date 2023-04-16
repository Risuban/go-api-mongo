package main

import (
	"context"
	a "mongoapi/api"
	"os"

	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/joho/godotenv"
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

	godotenv.Load()

	var server = os.Getenv("SERVER")
	var port = os.Getenv("PORT")
	var connection_string = os.Getenv("CONNECTION_STRING")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connection_string))
	if err != nil {
		panic(err)
	}
	fmt.Println(client)
	router := gin.Default()
	router.GET("/albums", a.GetAlbums)
	router.POST("/albums", a.PostAlbums)
	router.GET("/albums/:id", a.GetAlbumByID)
	addr := server + ":" + port
	router.Run(addr)
}
