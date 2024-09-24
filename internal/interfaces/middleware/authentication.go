package middleware

import (
	"hacktiv8-go-final-project/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)

		if err != nil {
			helpers.RespondWithError(c, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), err)
			return
		}

		c.Set("userData", verifyToken)
		c.Next()
	}
}
