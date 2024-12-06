package routes

import (
	"go-foodease-be/controller"
	"go-foodease-be/middleware"
	"go-foodease-be/service"

	"github.com/gin-gonic/gin"
)

func Order(route *gin.Engine, orderController controller.OrderController, jwtService service.JWTService) {
	routes := route.Group("/api/order")
	{
		routes.POST("/add", middleware.Authenticate(jwtService), orderController.AddToCart)

	}
}