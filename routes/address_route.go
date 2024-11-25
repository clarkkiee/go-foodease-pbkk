package routes

import (
	"go-foodease-be/controller"
	"go-foodease-be/middleware"
	"go-foodease-be/service"

	"github.com/gin-gonic/gin"
)

func Address(route *gin.Engine, addressController controller.AddressController, jwtService service.JWTService) {
	routes := route.Group("/api/address")
	{
		routes.POST("/new", middleware.Authenticate(jwtService) ,addressController.CreateNewAddress)
		routes.GET("/all", middleware.Authenticate(jwtService), addressController.GetAllAddressByCustomerId)
		routes.GET("/:address_id", middleware.Authenticate(jwtService), addressController.GetAdrressById)
		routes.PUT("/:address_id", middleware.Authenticate(jwtService), addressController.UpdateAddressById)
		routes.DELETE("/:address_id", middleware.Authenticate(jwtService), addressController.DeleteAddressById)
	}
}	