package main

import (
	"log"

	"github.com/ulisesaquiles1113/golmonrea/bd"
	"github.com/ulisesaquiles1113/golmonrea/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Handlers()
}
