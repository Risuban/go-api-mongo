package api

import (
	"context"
	"fmt"
	"os"
	"strconv"

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
	Numero_de_serie    int    `json:"numero_de_serie"`
	Stock_de_pasajeros int    `json:"stock_de_pasajeros"`
}
type ancillar struct {
	Nombre string `json:"nombre"`
	Stock  int    `json:"stock"`
	Ssr    string `json:"ssr"`
}
type vuelos struct {
	Vuelo        int        `json:"numero_vuelo"`
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

type payload_vuelo_put struct {
	Num_pas int `json:"stock_de_pasajeros"`
}

// getAlbums responds with the list of all albums as JSON.
func GetVuelos(c *gin.Context) {
	//traer información de la conexión
	godotenv.Load()
	var connection_string = os.Getenv("CONNECTION_STRING")
	// traer las variables para traer vuelos
	origen := c.Query("origen")
	destino := c.Query("destino")
	fecha := c.Query("fecha")
	filter := bson.M{
		"origen":  origen,
		"destino": destino,
		"fecha":   fecha,
	}
	//conexión
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connection_string))
	if err != nil {
		panic(err)
	}
	//Establecer conexión con collection
	usersCollection := client.Database("Tarea1").Collection("vuelos")

	cursor, err := usersCollection.Find(context.TODO(), filter)
	// chequear por errores
	if err != nil {
		panic(err)
	}

	// convert the cursor result to bson
	var results []bson.M
	// check for errors in the conversion
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	// display the documents retrieved
	fmt.Println("displaying all results from the search query")
	for _, result := range results {
		fmt.Println(result)
	}
	//aquí hay que ver por que el GET no QUIERE PARCEAR
	// codigo de la respuesta que recibe el cliente
	res := Response{}
	res.Code = 200
	res.Message = "Exito"
	res.Data = results

	c.JSON(res.Code, res)
	client.Disconnect(context.TODO())

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

	usersCollection := client.Database("Tarea1").Collection("reservas")
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

func PutVuelo(c *gin.Context) {
	var payload payload_vuelo_put
	c.ShouldBindJSON(&payload)
	numero_vuelo := c.Query("numero_vuelo")
	origen := c.Query("origen")
	destino := c.Query("destino")
	fecha := c.Query("fecha")

	godotenv.Load()
	var connection_string = os.Getenv("CONNECTION_STRING")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connection_string))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	usersCollection := client.Database("Tarea1").Collection("vuelos")

	numero_vuelo2, _ := strconv.Atoi(numero_vuelo)
	filter := bson.M{
		"numero_vuelo": numero_vuelo2,
		"origen":       origen,
		"destino":      destino,
		"fecha":        fecha,
	}

	update := bson.D{{"$set", bson.D{{"avion.stock_de_pasajeros", payload.Num_pas}}}}
	result, err := usersCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	var cursor vuelos
	problema := usersCollection.FindOne(context.TODO(), filter).Decode(&cursor)
	// chequear por errores
	if problema != nil {
		panic(problema)
	}
	fmt.Println(result)
	response := struct {
		Vuelo        int    `json:"numero_vuelo"`
		Origen       string `json:"origen"`
		Destino      string `json:"destino"`
		Hora_salida  string `json:"hora_salida"`
		Hora_llegada string `json:"hora_llegada"`
	}{
		Vuelo:        cursor.Vuelo,
		Origen:       cursor.Origen,
		Destino:      cursor.Destino,
		Hora_salida:  cursor.Hora_salida,
		Hora_llegada: cursor.Hora_llegada,
	}
	res := Response{}
	res.Code = 200
	res.Message = "Ok"
	res.Data = response
	c.JSON(res.Code, res)
	client.Disconnect(context.TODO())
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context) {

}
