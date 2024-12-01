package routes

import (
	"go-foodease-be/controller"
	"go-foodease-be/middleware"
	"go-foodease-be/service"

	"github.com/gin-gonic/gin"
)

func Product(route *gin.Engine, productController controller.ProductController, jwtService service.JWTService) {
    routes := route.Group("/api/product")
    {
        routes.POST("/create", middleware.Authenticate(jwtService), productController.CreateProduct)
        routes.PUT("/update/:product_id", middleware.Authenticate(jwtService), productController.UpdateProduct) 
    }
}

