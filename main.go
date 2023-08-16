package main

import (
	"crud-api/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Starting connection on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
