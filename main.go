package main

import (
	"log"
	"net/http"

	"github.com/fakh1m/go-auth-invoice/routes"
)

func main() {
	r := routes.SetupRouter()
	log.Println("Listening server : 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
