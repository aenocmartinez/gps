package domain

import (
	"abix/src/infraestructure/util"
	"log"
	"time"
)

type Person struct {
	id            int64
	name          string
	birthdate     string
	phone         string
	address       string
	email         string
	guardian      string
	guardianPhone string
	notes         []Note
	repository    PersonRepository
}

func NewPerson() *Person {
	return &Person{}
}

func (p *Person) SetRepository(repository PersonRepository) {
	p.repository = repository
}

func (p *Person) SetId(id int64) {
	p.id = id
}

func (p *Person) Id() int64 {
	return p.id
}

func (p *Person) SetName(name string) {
	p.name = name
}

func (p *Person) Name() string {
	return p.name
}

func (p *Person) SetBirthdate(birthdate string) {
	p.birthdate = birthdate
}

func (p *Person) Birthdate() string {
	return p.birthdate
}

func (p *Person) SetPhone(phone string) {
	p.phone = phone
}

func (p *Person) Phone() string {
	return p.phone
}

func (p *Person) SetAddress(address string) {
	p.address = address
}

func (p *Person) Address() string {
	return p.address
}

func (p *Person) SetEmail(email string) {
	p.email = email
}

func (p *Person) Email() string {
	return p.email
}

func (p *Person) SetGuardian(guardian string) {
	p.guardian = guardian
}

func (p *Person) Guardian() string {
	return p.guardian
}

func (p *Person) SetGuardianPhone(guardianPhone string) {
	p.guardianPhone = guardianPhone
}

func (p *Person) GuardianPhone() string {
	return p.guardianPhone
}

func (p *Person) Create() bool {
	return p.repository.CreatePerson(*p)
}

func (p *Person) Delete() bool {
	return p.repository.DeletePerson(p.id)
}

func (p *Person) Update() bool {
	return p.repository.UpdatePerson(*p)
}

func (p *Person) Age() int {

	now := time.Now()

	birthdate, err := time.Parse(util.YYYYmmdd, p.birthdate)
	if err != nil {
		log.Println(err)
		return 0
	}

	age := now.Year() - birthdate.Year()
	if now.Month() < birthdate.Month() ||
		(now.Month() == birthdate.Month() && now.Day() < birthdate.Day()) {
		age--
	}

	return age
}

func (p *Person) TodayIsYourBirthday() bool {
	now := time.Now()

	birthdate, err := time.Parse(util.YYYYmmdd, p.birthdate)
	if err != nil {
		log.Println(err)
		return false
	}

	return now.Equal(birthdate)
}

func (p *Person) AddNote(comment string, registerBy string) bool {
	p.notes = append(p.notes, *NewNote(comment, registerBy))
	return p.Update()
}

func (p *Person) RemoveNote(idNote int64) bool {
	var notes []Note = []Note{}
	for _, note := range p.notes {
		if note.id == idNote {
			continue
		}
		notes = append(notes, note)
	}
	p.notes = notes

	return p.Update()
}

func (p *Person) Notes() []Note {
	return p.notes
}

func (p *Person) Exists() bool {
	return p.id > 0
}
