package usecase

import "abix/src/view/dto"

type DeletePersonUseCase struct{}

func (uc *DeletePersonUseCase) Execute(id int64) (resp dto.Response) {

	person := personRepository.FindPersonById(id)
	if !person.Exists() {
		resp.Code = 400
		resp.Message = "resource not found"
		return resp
	}

	person.SetRepository(personRepository)
	success := person.Delete()
	if !success {
		resp.Code = 500
		resp.Message = "error server internal"
		return resp
	}

	resp.Code = 200
	resp.Message = "sucess"
	return resp
}
