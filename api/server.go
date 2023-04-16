package api

import (
	"github.com/gin-gonic/gin"
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

type destino struct {
	Vuelos    []vuelos   `json:"vuelos"`
	Pasajeros []pasajero `json:"pasajeros"`
}

// getAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {

}

// postAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {

}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context) {

}
