package main

import (
	"fmt"
	"github.com/ChrisDevOpsOrg/bookings/internal/handlers"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("application start listening on %s", portNumber))

	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		log.Fatal(err)
	}
}
