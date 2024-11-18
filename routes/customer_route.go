package routes

import (
	"go-foodease-be/controller"

	"github.com/gin-gonic/gin"
)

func Customer(route *gin.Engine, customerController controller.CustomerController) {
	routes := route.Group("/api/customer")
	{
		routes.GET("/me", customerController.Me)
		routes.POST("/login", customerController.Login)
	}
}