package dto

type PersonDto struct {
	Id            int64     `json:"id,omitempty"`
	Name          string    `json:"name,omitempty"`
	Birthdate     string    `json:"birthdate,omitempty"`
	Phone         string    `json:"phone,omitempty"`
	Address       string    `json:"address,omitempty"`
	Email         string    `json:"email,omitempty"`
	Guardian      string    `json:"guardian,omitempty"`
	GuardianPhone string    `json:"guardian_phone,omitempty"`
	Notes         []NoteDto `json:"notes,omitempty"`
}
