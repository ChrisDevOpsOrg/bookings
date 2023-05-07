package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("application start listening on %s", portNumber))

	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "about.page.tmpl")
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// check if the template in the cache
	_, inMap := tc[tmpl]
	if !inMap {
		log.Println("template cache not found")
		// create template cache
		CreateTemplateCache(tmpl)
	}

	t := tc[tmpl]
	fmt.Println("use template cache")
	t.Execute(w, nil)
}

// CreateTemplateCache create a template cache
var tc = map[string]*template.Template{}
var err error

func CreateTemplateCache(tmpl string) {
	templates := []string{
		fmt.Sprintf("./templates/%s", tmpl),
		"./templates/base.layout.tmpl",
	}

	tc[tmpl], err = template.ParseFiles(templates...)
	if err != nil {
		log.Fatal(err)
	}

}
