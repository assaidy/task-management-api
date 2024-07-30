package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/assaidy/task-manager-api/router"
)

const PORT = ":8080"

func main() {
	router := router.GetRouter()

	fmt.Printf("[INFO] Starting server at port %s...\n", PORT)

	log.Fatal(http.ListenAndServe(PORT, router))
}
