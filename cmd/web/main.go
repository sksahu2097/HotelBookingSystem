package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/sksahu2097/HotelBookingSystem/pkg/config"
	"github.com/sksahu2097/HotelBookingSystem/pkg/handlers"
	"github.com/sksahu2097/HotelBookingSystem/pkg/render"
)

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Error while creating application config")
	}
	app.TemplateCache = tc
	render.SetTemplateAppConfig(&app)
	repo := handlers.NewRepo(&app)
	handlers.SetRepo(repo)

	fmt.Println("Starting the application on 8080")
	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)

}
