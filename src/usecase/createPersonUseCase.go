package usecase

import "abix/src/view/dto"

type CreatePersonUseCase struct{}

func (uc *CreatePersonUseCase) Execute() (resp dto.Response) {
	resp.Code = 201
	resp.Message = "success"
	return resp
}
