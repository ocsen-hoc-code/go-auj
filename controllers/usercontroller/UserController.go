package usercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ocsen-hoc-code/go-auj/models/config"
	"github.com/ocsen-hoc-code/go-auj/models/user"
	userservice "github.com/ocsen-hoc-code/go-auj/services/usersevice"
)

type UserController struct {
	userServ  *userservice.UserService
	jwtConfig *config.JWTConfig
}

func NewUserController(userServInject *userservice.UserService, jwtConfigInject *config.JWTConfig) *UserController {
	return &UserController{userServ: userServInject, jwtConfig: jwtConfigInject}
}

func (ctrl UserController) Login(c *gin.Context) {

	user := user.User{}

	if err := c.ShouldBindJSON(&user); nil != err {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, errMsg := ctrl.userServ.Login(c.Request.Context(), user)
	if "" != errMsg {
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": token})
}
