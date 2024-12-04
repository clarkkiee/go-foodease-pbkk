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
    UpdateProduct(ctx *gin.Context) 
}

type productController struct {
	productService service.ProductService
	categoryService service.CategoryService
}

func NewProductController(ps service.ProductService, cs service.CategoryService) ProductController {
	return &productController{
		productService: ps,
		categoryService: cs,
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

	categoryId, err := c.categoryService.GetCategoryIdBySlug(ctx.Request.Context(), req.CategorySlug)
	if err != nil {
		response := utils.BuildFailedResponse("failed to identify category", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	createProcuct := dto.CreateProduct{
		ProductName: req.ProductName,
		Description: req.Description,
		PriceBefore: req.PriceBefore,
		PriceAfter: req.PriceAfter,
		ProductionTime: req.ProductionTime,
		ExpiredTime: req.ExpiredTime,
		Stock: req.Stock,
		CategoryID: categoryId,
		ImageID: req.ImageID,
	}

	res, err := c.productService.CreateProduct(ctx, createProcuct, storeID)
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

	categoryId, err := c.categoryService.GetCategoryIdBySlug(ctx.Request.Context(), req.CategorySlug)
	if err != nil {
		response := utils.BuildFailedResponse("failed to update product", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	req.CategorySlug = categoryId.String()

    res, err := c.productService.UpdateProduct(ctx, productID, req, storeID) 
    if err != nil {
        response := utils.BuildFailedResponse("failed to update product", err.Error(), nil)
        ctx.JSON(http.StatusBadRequest, response)
        return
    }

	response := utils.BuildSuccessResponse("product updated successfully", map[string]interface{}{"product_id": res})
    ctx.JSON(http.StatusOK, response)
}


