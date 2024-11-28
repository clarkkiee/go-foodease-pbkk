package controller

import (
	"go-foodease-be/dto"
	"go-foodease-be/service"
	"go-foodease-be/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController interface {
	CreateProduct(ctx *gin.Context)
}

type productController struct {
	productService service.ProductService
}

func NewProductController(ps service.ProductService) ProductController {
	return &productController{
		productService: ps,
	}
}

func (c *productController) CreateProduct(ctx *gin.Context) {
	storeID := ctx.MustGet("id").(string)
	var req dto.CreateProductRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response := utils.BuildFailedResponse("invalid request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	res, err := c.productService.CreateProduct(ctx, req, storeID)
	if err != nil {
		response := utils.BuildFailedResponse("failed to create product", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildSuccessResponse("product created successfully", res)
	ctx.JSON(http.StatusOK, response)
}
