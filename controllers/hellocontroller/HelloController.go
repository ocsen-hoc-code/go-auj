package hellocontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloController struct {
}

func NewHelloController() *HelloController {
	return &HelloController{}
}

func (ctrl HelloController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GIN Hello world"})
}
