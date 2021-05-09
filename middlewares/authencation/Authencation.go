package authencation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ocsen-hoc-code/go-auj/utils/jwtutil"
)

type Header struct {
	Authorization string `header:"Authorization" binding:"required"`
}

func Authentication(jwtSecretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := Header{}

		if err := c.ShouldBindHeader(&header); nil != err {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Missing token!"})
			return
		}

		user, ok := jwtutil.ValidToken(header.Authorization, jwtSecretKey)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Token is invalid!"})
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
