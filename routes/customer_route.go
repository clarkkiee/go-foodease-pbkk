package routes

import (
	"go-foodease-be/controller"
	"go-foodease-be/middleware"
	"go-foodease-be/service"

	"github.com/gin-gonic/gin"
)

func Customer(route *gin.Engine, customerController controller.CustomerController, jwtService service.JWTService) {
	routes := route.Group("/api/customer")
	{
		routes.GET("/me", customerController.Me)
		routes.POST("/login", customerController.Login)
		routes.POST("/register", customerController.Register)
		routes.DELETE("/delete", middleware.Authenticate(jwtService), customerController.DeleteAccount)
	}
}