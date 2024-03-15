package domain

import "abix/src/view/dto"

type PersonRepository interface {
	CreatePerson(person Person) bool
	UpdatePerson(person Person) bool
	DeletePerson(id int64) bool
	SearchPersons(criteria string) []Person
	FindPersonById(id int64) Person
	List() []dto.PersonDto
	BirthdayListMonth() []Person
	BirthdayListToday() []Person
}
