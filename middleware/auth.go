package middleware

import (
	"QueryEngine/helper"
	"QueryEngine/router/result"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		userClaim, err := helper.AnalyseToken(auth)

		if err != nil {
			c.Abort()
			c.JSON(http.StatusUnauthorized, result.Error("Unauthorized Authorization"))
			return
		}

		if userClaim == nil {
			c.Abort()
			c.JSON(http.StatusUnauthorized, result.Error("Unauthorized Admin"))
			return
		}
		c.Set("user_claims", userClaim)
		c.Next()
	}
}
