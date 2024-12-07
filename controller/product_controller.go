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

    DeleteProduct(ctx *gin.Context) // Menambahkan metode DeleteProduct
	GetProductById(ctx *gin.Context) 

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

	createProduct := dto.CreateProduct{
		ProductName:   req.ProductName,
		Description:   req.Description,
		PriceBefore:   req.PriceBefore,
		PriceAfter:    req.PriceAfter,
		ProductionTime: req.ProductionTime,
		ExpiredTime:   req.ExpiredTime,
		Stock:         req.Stock,
		CategoryID:    categoryId,
		ImageID:       req.ImageID,
	}

	res, err := c.productService.CreateProduct(ctx, createProduct, storeID)
	if err != nil {
		response := utils.BuildFailedResponse("failed to create product", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildSuccessResponse("product created successfully", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *productController) UpdateProduct(ctx *gin.Context) {
    productID := ctx.Param("product_id")  // Ambil product_id dari URL
    storeID := ctx.MustGet("id").(string) // Ambil store_id dari context JWT (auth)
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

func (c *productController) DeleteProduct(ctx *gin.Context) {
    productID := ctx.Param("product_id")  
    storeID := ctx.MustGet("id").(string)

    err := c.productService.DeleteProduct(ctx.Request.Context(), productID, storeID)
    if err != nil {
        response := utils.BuildFailedResponse("failed to delete product", err.Error(), nil)
        ctx.JSON(http.StatusBadRequest, response)
        return
    }

    response := utils.BuildSuccessResponse("product deleted successfully", nil)
    ctx.JSON(http.StatusOK, response)
}

func (c *productController) GetProductById(ctx *gin.Context) {
    productId := ctx.Param("product_id")
	res, err := c.productService.GetProductById(ctx.Request.Context(), productId)
	if err != nil {
		response := utils.BuildFailedResponse("Failed to get product", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return 
	}


	response := utils.BuildSuccessResponse("Get Product Successfully", res)
	ctx.JSON(http.StatusOK, response)
}
