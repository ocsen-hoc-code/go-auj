package hello2controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Hello2Controller struct {
}

func NewHello2Controller() *Hello2Controller {
	return &Hello2Controller{}
}

func (ctrl Hello2Controller) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to GIN 2"})
}
