package absence

import (
	"html"
	"log"
	"strconv"
	"time"
)

type Absence struct {
	Id                      string    `json:"id,omitempty"`
	EmployeeId              string    `json:"employee_id,omitempty"`
	Kind                    int       `json:"kind,omitempty"`
	From                    time.Time `json:"from,omitempty"`
	To                      time.Time `json:"to,omitempty"`
	Reason                  string    `json:"reason,omitempty"`
	ClarificationMp         bool      `json:"clarification_mp,omitempty"`
	ClarificationSubstitude bool      `json:"clarification_substitude,omitempty"`
	Approved                bool      `json:"approved,omitempty"`
	Comment                 string    `json:"comment,omitempty"`
	WorkdaysPerWeek         int       `json:workdays_per_week,omitempty`
	CreatedAt               time.Time `json:"created_at,omitempty"`
}

func Create(
	//employeeId string,
	kind string,
	from string,
	to string,
	reason string,
	clarificationMp string,
	clarificationSubstitude string,
	comment string,
	workdaysPerWeek string) (absence Absence, errArr []error) {

	var (
		errFrom                    error
		errTo                      error
		errKind                    error
		errClarificationMp         error
		errClarificationSubstitude error
		errWorkdaysPerWeek         error
	)

	log.Printf("Kind: %+v", kind)
	log.Printf("From: %+v", from)
	log.Printf("To: %+v", to)
	log.Printf("Reason: %+v", reason)
	log.Printf("ClarificationMp: %+v", clarificationMp)
	log.Printf("ClarificationSubstitude: %+v", clarificationSubstitude)
	log.Printf("Comment: %+v", comment)
	log.Printf("WorkdaysPerWeek: %+v", workdaysPerWeek)

	layout := "2006-01-02"

	//absence := Absence{}

	absence.Kind, errKind = strconv.Atoi(kind)

	if errKind != nil {
		log.Println(errKind.Error())
		errArr = append(errArr, errKind)
	}

	absence.From, errFrom = time.Parse(layout, from)

	if errFrom != nil {
		log.Println(errFrom.Error())
		errArr = append(errArr, errFrom)
	}

	absence.To, errTo = time.Parse(layout, to)

	if errTo != nil {
		log.Println(errTo.Error())
		errArr = append(errArr, errTo)
	}

	absence.Reason = html.EscapeString(reason)

	absence.ClarificationMp, errClarificationMp = strconv.ParseBool(clarificationMp)

	if errClarificationMp != nil {
		log.Println(errClarificationMp.Error())
		errArr = append(errArr, errClarificationMp)
	}

	absence.ClarificationSubstitude, errClarificationSubstitude = strconv.ParseBool(clarificationSubstitude)

	if errClarificationSubstitude != nil {
		log.Println(errClarificationSubstitude.Error())
		errArr = append(errArr, errClarificationSubstitude)
	}

	absence.Comment = html.EscapeString(comment)

	absence.WorkdaysPerWeek, errWorkdaysPerWeek = strconv.Atoi(workdaysPerWeek)

	if errWorkdaysPerWeek != nil {
		log.Println(errWorkdaysPerWeek.Error())
		errArr = append(errArr, errWorkdaysPerWeek)
	}

	absence.CreatedAt = time.Now()

	log.Printf("errArr: %+v", errArr)

	if errArr != nil {
		return
	}

	GeneratePdf(absence)

	return
}

func Get(employeeId string) {

}
