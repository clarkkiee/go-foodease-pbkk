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
        routes.DELETE("/delete/:product_id", middleware.Authenticate(jwtService), productController.DeleteProduct) 
        routes.POST("/create", middleware.Authenticate(jwtService), productController.CreateProduct)
        routes.PUT("/update/:product_id", middleware.Authenticate(jwtService), productController.UpdateProduct)
        routes.GET("/:product_id", middleware.Authenticate(jwtService), productController.GetProductById)
        routes.GET("/store", middleware.Authenticate(jwtService), productController.GetProductByStoreId)
        routes.GET("/customer/:customer_id?limit=20&offset=0&distance=20000", middleware.Authenticate(jwtService), productController.GetNearestProduct)
    }

}
