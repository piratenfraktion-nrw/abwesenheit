package absence

type Absence struct {
	Id                      string `json:"id,omitempty"`
	EmployeeId              string `json:"employee_id,omitempty"`
	Kind                    string `json:"kind,omitempty"`
	From                    string `json:"from,omitempty"`
	Reason                  string `json:"reason,omitempty"`
	ClarificationMp         string `json:"clarification_mp,omitempty"`
	ClarificationSubstitude string `json:"clarification_substitude,omitempty"`
	Approved                string `json:"approved,omitempty"`
	Comment                 string `json:"comment,omitempty"`
	CreatedAt               string `json:"created_at,omitempty"`
}

func Create(employeeId string) {

}
