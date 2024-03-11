package domain

type PersonRepository interface {
	CreatePerson(person Person) bool
	UpdatePerson(person Person) bool
	DeletePerson(id int64) bool
	SearchPersons(criteria string) []Person
	FindPersonById(id int64) Person
	List() []Person
	BirthdayListMonth() []Person
	BirthdayListToday() []Person
}
