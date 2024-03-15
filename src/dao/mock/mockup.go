package mock

import (
	"abix/src/domain"
	"abix/src/view/dto"
)

type Mockup struct {
	databases map[int64]domain.Person
}

func NewMockup() *Mockup {
	return &Mockup{
		databases: make(map[int64]domain.Person),
	}
}

func (m *Mockup) CreatePerson(person domain.Person) bool {
	nextCurrentId := int64(len(m.databases)) + 1

	person.SetId(nextCurrentId)
	m.databases[nextCurrentId] = person
	return true
}

func (m *Mockup) UpdatePerson(person domain.Person) bool {
	return true
}

func (m *Mockup) DeletePerson(id int64) bool {
	return true
}

func (m *Mockup) SearchPersons(criteria string) (persons []domain.Person) {
	return persons
}

func (m *Mockup) FindPersonById(id int64) (person domain.Person) {
	person, exists := m.databases[id]
	if !exists {
		return domain.Person{}
	}
	return person
}

func (m *Mockup) List() (persons []dto.PersonDto) {
	persons = []dto.PersonDto{}
	for _, person := range m.databases {
		persons = append(persons, dto.PersonDto{
			Id:            person.Id(),
			Name:          person.Name(),
			Birthdate:     person.Birthdate(),
			Phone:         person.Phone(),
			Address:       person.Address(),
			Email:         person.Email(),
			Guardian:      person.Guardian(),
			GuardianPhone: person.GuardianPhone(),
		})
	}

	return persons
}

func (m *Mockup) BirthdayListMonth() (persons []domain.Person) {
	return persons
}

func (m *Mockup) BirthdayListToday() (persons []domain.Person) {
	return persons
}
