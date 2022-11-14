package middleware

import (
	"fmt"
	"log"
	"net/http"
	"pustaka-api/helper"
	"pustaka-api/jwtUse"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AuthorizeJWT validates the token user given, return 401 if not valid
func AuthorizeJWT(jwtService jwtUse.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		// authHeader := c.GetHeader("Authorization")
		Get := session.Get("token")
		authHeader := fmt.Sprintf("%v", Get)
		if Get == nil {
			response := helper.BuildErrorResponse("Failed to process request", "No token found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtUse.NewJWTService().ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]: ", claims["user_id"])
			log.Println("Claim[issuer] :", claims["issuer"])
		} else {
			log.Println(err)
			response := helper.BuildErrorResponse("Token is not valid", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
