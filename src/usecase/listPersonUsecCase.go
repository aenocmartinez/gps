package usecase

import "abix/src/view/dto"

type ListPersonUsecCase struct{}

func (uc *ListPersonUsecCase) Execute() (resp dto.Response) {

	resp.Code = 200
	resp.Data = personRepository.List()

	return resp
}
