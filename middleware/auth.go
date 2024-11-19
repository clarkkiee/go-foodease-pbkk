package middleware

import (
	"go-foodease-be/service"
	"go-foodease-be/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate(jwtService service.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			response := utils.BuildFailedResponse("Failed processing request", "Token not found", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		if !strings.Contains(authHeader, "Bearer ") {
			response := utils.BuildFailedResponse("Failed processing request", "Invalid Auth Method", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		authHeader = strings.Replace(authHeader, "Bearer ", "", -1)
		token, err := jwtService.ValidateToken(authHeader)
		if err != nil {
			response := utils.BuildFailedResponse("Failed processing request", "Invalid token", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		if !token.Valid {
			response := utils.BuildFailedResponse("Failed processing request", "Denied Access", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		id, err := jwtService.GetEntityIdByToken(authHeader)
		if err != nil {
			response := utils.BuildFailedResponse("Failed processing request", err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		
		ctx.Set("id", id)
		ctx.Set("token", authHeader)
		ctx.Next()
	}
}