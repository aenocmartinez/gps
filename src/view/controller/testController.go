package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestMutant(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}
