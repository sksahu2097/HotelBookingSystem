package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/sksahu2097/HotelBookingSystem/pkg/config"
	"github.com/sksahu2097/HotelBookingSystem/pkg/models"
)

// Render Template using html template
func RenderTenplateTe(w http.ResponseWriter, tmpl string) {
	parseTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parseTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing the template ", err)
		return
	}
}

var app *config.AppConfig

// setTemplateAppconfig to resduce the boilerPlateCode
func SetTemplateAppConfig(a *config.AppConfig) {
	app = a
}

func RenderTenplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	//create the template cache
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	tc = app.TemplateCache

	//get requested template from cache
	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("Template not found in the Templates cache")
	}

	buff := new(bytes.Buffer)

	err := t.Execute(buff, td)
	if err != nil {
		log.Println(err)
	}

	_, err = buff.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := make(map[string]*template.Template)

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	//range over each file and add to cache
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts

	}
	return myCache, nil

}
