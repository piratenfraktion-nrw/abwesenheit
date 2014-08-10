package main

import (
	"github.com/gorilla/mux"
	"github.com/piratenfraktion-nrw/abwesenheit/api"
	"github.com/piratenfraktion-nrw/abwesenheit/configuration"
	"log"
	"net/http"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU()) // use all CPU cores

	err := configuration.Load() // TODO: Schauen, ob es sinnig ist, ein Load() zu Beginn von Save() aufzurufen, damit vorhande Werte nicht geloescht werden
	configuration.Config.Port = ":8080"
	configuration.Save()

	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Starting server...")

	m := mux.NewRouter()

	m.HandleFunc("/absence/create", api.CreateAbsence).Methods("POST")

	// Everything else fails
	m.HandleFunc("/{path:.*}", http.NotFound)

	log.Println("Now listening on port", configuration.Config.Port)

	http.Handle("/", m)
	http.ListenAndServe(configuration.Config.Port, nil)
}
