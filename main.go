package main

import (
	a "mongoapi/api"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

// definir bien puerto
func main() {

	godotenv.Load()

	var server = os.Getenv("SERVER")
	var port = os.Getenv("PORT")

	router := gin.Default()
	router.GET("/api/vuelo", a.GetVuelos)
	router.POST("/api/reserva", a.PostReserva)
	router.GET("/albums/:id", a.GetAlbumByID)
	addr := server + ":" + port
	router.Run(addr)

}
