package controller

import (
	"go-foodease-be/dto"
	"go-foodease-be/service"
	"go-foodease-be/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)
import "fmt"
type ProductController interface {
    CreateProduct(ctx *gin.Context)
    UpdateProduct(ctx *gin.Context) 
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





func (c *productController) UpdateProduct(ctx *gin.Context) {
    productID := ctx.Param("product_id")  
    storeID := ctx.MustGet("id").(string) 
    var req dto.UpdateProductRequest
    if err := ctx.ShouldBind(&req); err != nil {
        response := utils.BuildFailedResponse("invalid request", err.Error(), nil)
        ctx.JSON(http.StatusBadRequest, response)
        return
    }
	fmt.Printf("PriceBefore: %v (type: %T)\n", req.PriceBefore, req.PriceBefore)
    fmt.Printf("PriceAfter: %v (type: %T)\n", req.PriceAfter, req.PriceAfter)

    res, err := c.productService.UpdateProduct(ctx, productID, req, storeID) 
    if err != nil {
        response := utils.BuildFailedResponse("failed to update product", err.Error(), nil)
        ctx.JSON(http.StatusBadRequest, response)
        return
    }

    response := utils.BuildSuccessResponse("product updated successfully", res)
    ctx.JSON(http.StatusOK, response)
}

