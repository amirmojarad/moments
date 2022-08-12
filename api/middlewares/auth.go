package middlewares

import (
	"github.com/gin-gonic/gin"
	"moments/api/auth"
	"net/http"
	"strings"
)

func CheckAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwtService := auth.New()
		header := ctx.Request.Header
		if authToken, ok := header["Authorization"]; ok {
			token := strings.Split(authToken[0], " ")[1]
			if t, err := jwtService.ValidateToken(token); err != nil {
				ctx.AbortWithError(http.StatusUnauthorized, err)
				return
			} else {
				claims := jwtService.GetMapClaims(t)
				ctx.Set("username", claims["username"])
				ctx.Next()
			}
		} else {
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{
				"error": "request does not contain any token in header",
			})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
