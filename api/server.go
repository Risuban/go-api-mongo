package api

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/joho/godotenv"
)

type avion struct {
	Modelo             string `json:"modelo"`
	Numero_de_serie    string `json:"numero_de_serie"`
	Stock_de_pasajeros string `json:"stock_de_pasajeros"`
}
type ancillar struct {
	nombre string `json:"nombre"`
	stock  string `json:"stock"`
	ssr    string `json:"ssr"`
}
type vuelos struct {
	Vuelo        string     `json:"numero_vuelo"`
	Origen       string     `json:"origen"`
	Destino      string     `json: "destino" `
	Hora_salida  string     `json:"hora_salida"`
	Hora_llegada string     `json:"hora_llegada"`
	Avion        avion      `json:"avion"`
	Ancillaries  []ancillar `json:"ancillaries"`
}
type viaje struct {
	Ssr      string `json:"ssr"`
	Cantidad string `json:"cantidad"`
}
type ancillaries_pasajero struct {
	Ida    []viaje `json:"ida"`
	Vuelta []viaje `json:"vuelta"`
}
type balances struct {
	Ancillaries_ida    int `json:"Ancillaries_ida"`
	vuelo_ida          int `json:"vuelo_ida"`
	ancillaries_vuelta int `json:"ancillaries_vuelta"`
	vuelo_vuelta       int `json:"vuelo_vuelta"`
}
type pasajero struct {
	Nombre      string               `json:"nombre"`
	Apellido    string               `json:"apellido"`
	Edad        int                  `json:"edad"`
	Ancillaries ancillaries_pasajero `json:"ancillaries"`
	Balances    balances             `json:"balances"`
}

type reserva struct {
	Vuelos    []vuelos   `json:"vuelos"`
	Pasajeros []pasajero `json:"pasajeros"`
}

// getAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {

}

// postAlbums adds an album from JSON received in the request body.
func PostReserva(c *gin.Context) {
	//recepción del payload
	var payload avion
	c.ShouldBindJSON(&payload)
	fmt.Printf(payload.Modelo)
	fmt.Printf(payload.Numero_de_serie)
	fmt.Printf(payload.Stock_de_pasajeros)
	//out, err := bson.MarshalExtJSON(payload, false, false)
	//detalles de la conexión con el mongo
	godotenv.Load()
	var connection_string = os.Getenv("CONNECTION_STRING")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connection_string))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	usersCollection := client.Database("testing").Collection("users")
	user := bson.D{{"modelo", payload.Modelo}, {"numero_de_serie", payload.Numero_de_serie}}
	// insert the bson object using InsertOne()
	result, err := usersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}
	// display the id of the newly inserted object
	fmt.Println(result.InsertedID)

}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context) {

}
