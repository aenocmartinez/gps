package usecase

import "abix/src/view/dto"

type FindUserByIdUseCase struct{}

func (uc *FindUserByIdUseCase) Execute(personId int64) (resp dto.Response) {

	person := personRepository.FindPersonById(personId)

	if !person.Exists() {
		resp.Code = 404
		resp.Message = "resource not found"
		return resp
	}

	resp.Code = 200
	resp.Data = dto.PersonDto{
		Id:            person.Id(),
		Name:          person.Name(),
		Birthdate:     person.Birthdate(),
		Phone:         person.Phone(),
		Address:       person.Address(),
		Email:         person.Email(),
		Guardian:      person.Guardian(),
		GuardianPhone: person.GuardianPhone(),
	}

	return resp
}
