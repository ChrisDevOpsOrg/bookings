package main

import (
	"fmt"
	"github.com/ChrisDevOpsOrg/bookings/internal/config"
	"github.com/ChrisDevOpsOrg/bookings/internal/handlers"
	"github.com/ChrisDevOpsOrg/bookings/internal/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

var app config.AppConfig

func main() {

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache", err)
	}

	app.TemplateCache = tc

	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("application start listening on %s", portNumber))

	err = http.ListenAndServe(portNumber, nil)
	if err != nil {
		log.Fatal(err)
	}
}
