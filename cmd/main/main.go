package main

import (
	"log"
	"net/http"

	"github.com/imchukwu/school/pkg/routes"
)

func main() {
	log.Println("Starting the server...")
	router := routes.SetupRouter()
	log.Println("Server is listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
