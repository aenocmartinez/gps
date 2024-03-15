package main

import (
	"abix/src/view/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/mutant", controller.TestMutant)

	r.GET("/persons", controller.ListPersons)
	r.GET("/persons/:id", controller.FindPersonById)
	r.POST("/persons", controller.CreatePerson)
	r.PUT("/persons", controller.UpdatePerson)
	r.DELETE("/persons/:id", controller.DeletePerson)

	r.Run(":8081")
}
