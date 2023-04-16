package main

import (
	"context"
	a "mongoapi/api"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/joho/godotenv"
)

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
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	router := gin.Default()
	router.GET("/albums", a.GetAlbums)
	router.POST("/albums", a.PostAlbums)
	router.GET("/albums/:id", a.GetAlbumByID)
	addr := server + ":" + port
	router.Run(addr)

}
