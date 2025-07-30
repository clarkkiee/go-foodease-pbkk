package main

import (
	"log"
	"os"

	"go-foodease-be/config"
	"go-foodease-be/controller"
	"go-foodease-be/middleware"
	"go-foodease-be/pkg/logger"
	"go-foodease-be/pkg/metrics"
	"go-foodease-be/repository"
	"go-foodease-be/routes"
	"go-foodease-be/service"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	db := config.DatabaseConnection()
	logger.Initialize()
	defer logger.Sync()

	metrics.Initialize()

	var (
		jwtService service.JWTService = service.NewJWTService()
		customerRepository repository.CustomerRepository = repository.NewCustomerRepository(db)
		customerService service.CustomerService = service.NewCustomerService(customerRepository, jwtService)
		customerController controller.CustomerController = controller.NewCustomerController(customerService)

		addressRepository repository.AddressRepository = repository.NewAddressRepository(db)
		addressService service.AddressService = service.NewAddressService(addressRepository, jwtService)
		addressController controller.AddressController = controller.NewAddressController(addressService)

		categoryRepository repository.CategoryRepository = repository.NewCategoryRepository(db)
		categoryService service.CategoryService = service.NewCategoryService(categoryRepository)

		productRepository repository.ProductRepository = repository.NewProductRepository(db)
		productService service.ProductService = service.NewProductService(productRepository, addressRepository, jwtService)
		productController controller.ProductController = controller.NewProductController(productService, categoryService, addressService)
		
		storeRepository repository.StoreRepository = repository.NewStoreRepository(db)
		storeService service.StoreService = service.NewStoreService(storeRepository, jwtService)
		storeController controller.StoreController = controller.NewStoreController(storeService, addressService)

		orderRepository repository.OrderRepository = repository.NewOrderRepository(db)
		orderService service.OrderService = service.NewOrderService(orderRepository, productRepository, db)
		orderController controller.OrderController = controller.NewOrderController(orderService, productService)
	)

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())
	server.Use(middleware.LoggerMiddleware())
	server.Use(middleware.MetricsMiddleware())

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "fong",
		})
	})

	server.GET("/success", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Success"})
	})
	server.GET("/client-error", func(c *gin.Context) {
		c.JSON(400, gin.H{"error": "Client error"})
	})
	server.GET("/server-error", func(c *gin.Context) {
		c.JSON(500, gin.H{"error": "Server error"})
	})

	server.GET("/metrics", gin.WrapH(promhttp.Handler()))

	routes.Customer(server, customerController, jwtService)
	routes.Address(server,addressController, jwtService)
	routes.Product(server, productController, jwtService)
	routes.Store(server, storeController, jwtService)
	routes.Order(server, orderController, jwtService)

	server.Static("/assets", "./assets")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	var serve string
	if os.Getenv("APP_ENV") == "localhost" {
		serve = "127.0.0.1:" + port
	} else {
		serve = ":" + port
	}

	if err := server.Run(serve); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}