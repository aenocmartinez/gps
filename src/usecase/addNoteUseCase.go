package usecase

import "abix/src/view/dto"

type AddNoteUseCase struct{}

func (uc *AddNoteUseCase) Execute(personId int64, comment string, registerBy string) (resp dto.Response) {

	person := personRepository.FindPersonById(personId)

	if !person.Exists() {
		resp.Code = 404
		resp.Message = "resource not found"
		return resp
	}

	person.SetRepository(personRepository)
	success := person.AddNote(comment, registerBy)
	if !success {
		resp.Code = 500
		resp.Message = "server error internal"
	}

	resp.Code = 201
	resp.Message = "success"
	return resp
}
