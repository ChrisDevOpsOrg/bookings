package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// call CreateTemplateCache function to create template cache
	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal("failed to create template cache")
	}

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("cannot find template in cache", err)
	}

	buf := new(bytes.Buffer)

	err = t.Execute(buf, nil)
	if err != nil {
		log.Fatal(err)
	}

	// render template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{}

	// get all the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return templateCache, err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return templateCache, err
		}

		// get all layout templates
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return templateCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
		}

		templateCache[name] = ts
	}

	return templateCache, nil
}
