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

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type avion struct {
	Modelo             string `json:"modelo"`
	Numero_de_serie    string `json:"numero_de_serie"`
	Stock_de_pasajeros string `json:"stock_de_pasajeros"`
}
type ancillar struct {
	Nombre string `json:"nombre"`
	Stock  string `json:"stock"`
	Ssr    string `json:"ssr"`
}
type vuelos struct {
	Vuelo        string     `json:"numero_vuelo"`
	Origen       string     `json:"origen"`
	Destino      string     `json:"destino" `
	Hora_salida  string     `json:"hora_salida"`
	Hora_llegada string     `json:"hora_llegada"`
	Fecha        string     `json:"fecha"`
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
	Vuelo_ida          int `json:"vuelo_ida"`
	Ancillaries_vuelta int `json:"ancillaries_vuelta"`
	Vuelo_vuelta       int `json:"vuelo_vuelta"`
}
type pasajero struct {
	Nombre      string                 `json:"nombre"`
	Apellido    string                 `json:"apellido"`
	Edad        int                    `json:"edad"`
	Ancillaries []ancillaries_pasajero `json:"ancillaries"`
	Balances    balances               `json:"balances"`
}

type reserva struct {
	Vuelos    []vuelos   `json:"vuelos"`
	Pasajeros []pasajero `json:"pasajeros"`
}

// getAlbums responds with the list of all albums as JSON.
func GetVuelos(c *gin.Context) {

	// Retrieve the values of the "origen", "destino", and "fecha" query parameters
	origen := c.Query("origen")
	destino := c.Query("destino")
	fecha := c.Query("fecha")
	print(destino, fecha)
	/*if err != nil {
		res := Response{}
		res.Code = 200
		res.Message = "Error al procesar los parametros"

		c.JSON(res.Code, res)
	}*/
	//aquí hay que ver por que el GET no QUIERE PARCEAR
	//uwu := info_vuelo.GET("origen")
	// codigo de la respuesta que recibe el cliente
	res := Response{}
	res.Code = 200
	res.Message = origen

	c.JSON(res.Code, res)

}

// postAlbums adds an album from JSON received in the request body.
func PostReserva(c *gin.Context) {
	//recepción del payload
	var payload reserva
	c.ShouldBindJSON(&payload)

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
	user := bson.D{{"vuelos", payload.Vuelos}, {"pasajeros", payload.Pasajeros}}
	// insert the bson object using InsertOne()
	result, err := usersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}
	// codigo de la respuesta que recibe el cliente
	fmt.Println(result.InsertedID)
	res := Response{}
	res.Code = 200
	res.Message = "Ok"

	c.JSON(res.Code, res)
	client.Disconnect(context.TODO())

}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context) {

}
