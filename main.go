package main

import (
	"abix/src/view/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/mutant", controller.TestMutant)

	r.Run(":8081")
}
