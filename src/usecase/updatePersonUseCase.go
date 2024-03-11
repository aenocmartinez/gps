package usecase

import "abix/src/view/dto"

type UpdatePersonUseCase struct{}

func (uc *UpdatePersonUseCase) Execute(personDto dto.PersonDto) (resp dto.Response) {

	person := personRepository.FindPersonById(personDto.Id)
	if !person.Exists() {
		resp.Code = 404
		resp.Message = "resource not found"
		return resp
	}

	person.SetId(personDto.Id)
	person.SetName(personDto.Name)
	person.SetPhone(personDto.Phone)
	person.SetAddress(personDto.Address)
	person.SetBirthdate(personDto.Birthdate)
	person.SetEmail(personDto.Email)
	person.SetGuardian(personDto.Guardian)
	person.SetGuardianPhone(personDto.GuardianPhone)
	person.SetRepository(personRepository)

	success := person.Update()
	if !success {
		resp.Code = 500
		resp.Message = "server error internal"
		return resp
	}

	resp.Code = 200
	resp.Message = "success"
	return resp
}
