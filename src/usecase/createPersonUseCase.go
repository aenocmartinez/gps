package usecase

import (
	"abix/src/domain"
	"abix/src/view/dto"
)

type CreatePersonUseCase struct{}

func (uc *CreatePersonUseCase) Execute(personDto dto.PersonDto) (resp dto.Response) {

	person := domain.NewPerson()
	person.SetRepository(personRepository)
	person.SetName(personDto.Name)
	person.SetAddress(personDto.Address)
	person.SetBirthdate(personDto.Birthdate)
	person.SetEmail(personDto.Email)
	person.SetGuardian(personDto.Guardian)
	person.SetGuardianPhone(personDto.GuardianPhone)
	person.SetPhone(personDto.Phone)

	success := person.Create()
	if !success {
		resp.Code = 500
		resp.Message = "error internal server"
		return resp
	}

	resp.Code = 201
	resp.Message = "success"
	return resp
}
