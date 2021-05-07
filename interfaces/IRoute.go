package interfaces

import "github.com/gin-gonic/gin"

type IRoute interface {
	RouteRegister(*gin.Engine)
}
