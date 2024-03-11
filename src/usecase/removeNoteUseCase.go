package usecase

import "abix/src/view/dto"

type RemoveNoteUseCase struct{}

func (uc *RemoveNoteUseCase) Execute(personId, noteId int64) (resp dto.Response) {

	person := personRepository.FindPersonById(personId)
	if !person.Exists() {
		resp.Code = 404
		resp.Message = "resource not found"
		return resp
	}

	person.SetRepository(personRepository)
	success := person.RemoveNote(noteId)
	if !success {
		resp.Code = 500
		resp.Message = "error server internal"
		return resp
	}

	resp.Code = 200
	resp.Message = "success"
	return resp
}
