package api

import (
	"encoding/json"
	"github.com/piratenfraktion-nrw/abwesenheit/absence"
	"log"
	"net/http"
	"time"
)

type Result struct {
	Result interface{} `json:",omitempty"`
	Error  []error     `json:",omitempty"`
	Time   time.Time   `json:",omitempty"`
}

func CreateAbsence(w http.ResponseWriter, r *http.Request) {

	log.Println("Create new absence entry...")

	var result Result

	r.ParseForm()
	absence, errArr := absence.Create(
		//r.PostFormValue("employee_id"),
		r.PostFormValue("kind"),
		r.PostFormValue("from"),
		r.PostFormValue("to"),
		r.PostFormValue("reason"),
		r.PostFormValue("clarification_mp"),
		r.PostFormValue("clarification_substitude"),
		r.PostFormValue("comment"),
		r.PostFormValue("workdays_per_week"))

	result.Time = time.Now()

	if errArr != nil {
		result.Error = errArr
		result.Time = time.Now()
	} else {
		result.Result = absence
	}

	log.Printf("Result: %+v", result)

	res, errRes := json.Marshal(result)

	if errRes != nil {
		w.Write([]byte("{ Error: \"Unknown error\", Time: \"" + time.Now().String() + "\"}"))
	} else {
		w.Write(res)
	}
}

func GetAbsence(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	log.Println("Get absence entries for employee_id:", r.PostFormValue("employee_id"))

	absence.Get(r.PostFormValue("employee_id"))
}
