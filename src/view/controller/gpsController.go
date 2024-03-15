package controller

import (
	"abix/src/infraestructure/util"
	"abix/src/usecase"
	"abix/src/view/dto"
	formrequest "abix/src/view/form-request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListPersons(c *gin.Context) {

	listPersons := usecase.ListPersonUsecCase{}
	c.JSON(http.StatusOK, listPersons.Execute())
}

func CreatePerson(c *gin.Context) {
	var req formrequest.CreatePersonFormRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	createPerson := usecase.CreatePersonUseCase{}
	response := createPerson.Execute(dto.PersonDto{
		Name:          req.Name,
		Birthdate:     req.Birthdate,
		Phone:         req.Phone,
		Email:         req.Email,
		Address:       req.Address,
		Guardian:      req.Guardian,
		GuardianPhone: req.GuardianPhone,
	})

	c.JSON(response.Code, response)
}

func UpdatePerson(c *gin.Context) {
	var req formrequest.UpdatePersonFormRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	updatePerson := usecase.UpdatePersonUseCase{}
	response := updatePerson.Execute(dto.PersonDto{
		Name:          req.Name,
		Birthdate:     req.Birthdate,
		Phone:         req.Phone,
		Email:         req.Email,
		Address:       req.Address,
		Guardian:      req.Guardian,
		GuardianPhone: req.GuardianPhone,
	})

	c.JSON(response.Code, response.Message)
}

func FindPersonById(c *gin.Context) {

	id, err := util.ConvertStringToInt64(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "param invalid"})
		return
	}

	findPerson := usecase.FindUserByIdUseCase{}
	response := findPerson.Execute(id)
	c.JSON(response.Code, response)
}

func DeletePerson(c *gin.Context) {
	id, err := util.ConvertStringToInt64(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "param invalid"})
		return
	}

	deletePerson := usecase.DeletePersonUseCase{}
	response := deletePerson.Execute(id)
	c.JSON(response.Code, response)
}
