package routes

import (
	"go-foodease-be/controller"
	// "go-foodease-be/middleware"
	"go-foodease-be/service"

	"github.com/gin-gonic/gin"
)

func Store(route *gin.Engine, storeController controller.StoreController, jwtService service.JWTService) {
	routes := route.Group("/api/store")
	{
		routes.POST("/login", storeController.Login)
	}
}