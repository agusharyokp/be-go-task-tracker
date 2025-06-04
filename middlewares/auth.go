package middlewares

import (
	"net/http"
	"strings"

	"github.com/yourname/go-task-tracker/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	authHeader := context.Request.Header.Get("Authorization")
	if authHeader == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message":"Unauthorized !"})
		return 
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid Authorization header format"})
		return
	}
	token := tokenParts[1]

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message":err.Error()})
		return
	}

	context.Set("userId", userId)
	context.Next()
}