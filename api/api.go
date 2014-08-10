package api

import (
	"github.com/piratenfraktion-nrw/abwesenheit/absence"
	"log"
	"net/http"
)

func CreateAbsence(w http.ResponseWriter, r *http.Request) {

	log.Println("Create new absence entry...")

	r.ParseForm()
	absence.Create(r.PostFormValue("1"))
}
