package middleware

import (
	"gin-socmed/errorhandler"
	"gin-socmed/helper"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			errorhandler.HandlerError(c, &errorhandler.UnauthorizedError{Message: "Unauthorize"})
			c.Abort()
			return
		}

		userId, err := helper.ValidateToken(tokenString)
		if err != nil {
			errorhandler.HandlerError(c, &errorhandler.UnauthorizedError{Message: err.Error()})
			c.Abort()
			return
		}

		c.Set("userID", *userId)
		c.Next()
	}
}
