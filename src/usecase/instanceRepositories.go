package usecase

import (
	"abix/src/dao/mock"
	"abix/src/domain"
)

var personRepository domain.PersonRepository = mock.NewMockup()
