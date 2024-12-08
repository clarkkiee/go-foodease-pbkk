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
	GetProductByStoreId(ctx *gin.Context)
	GetNearestProduct(ctx *gin.Context)
}

type productController struct {
	productService service.ProductService
	categoryService service.CategoryService
	addressService service.AddressService
}

func NewProductController(ps service.ProductService, cs service.CategoryService, as service.AddressService) ProductController {
	return &productController{
		productService: ps,
		categoryService: cs,
		addressService: as,
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

func (c *productController) GetProductByStoreId(ctx *gin.Context) {
	storeId := ctx.MustGet("id").(string)
	res, err := c.productService.GetProductByStoreId(ctx.Request.Context(), storeId)
	if err != nil {
		response := utils.BuildFailedResponse("Failed to get product", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return 
	}

	response := utils.BuildSuccessResponse("Get Product Successfully", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *productController) GetNearestProduct(ctx *gin.Context) {
	customerId := ctx.MustGet("id").(string)
	limit := ctx.Query("limit")
	offset := ctx.Query("offset")
	maxDistance := ctx.Query("distance")

	customerAddress, err := c.addressService.GetActiveAddress(ctx.Request.Context(), customerId)
	if err != nil {
		response := utils.BuildFailedResponse("Failed to get customer location", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return 
	}
	
	customerCoord := dto.CoordinatesResponse{
		Longitude: customerAddress.Longitude,
		Latitude: customerAddress.Latitude,
	}

	res, err := c.productService.GetNearestProduct(ctx.Request.Context(), customerCoord, limit, offset, maxDistance)
	if err != nil {
		response := utils.BuildFailedResponse("Failed to get product", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return 
	}

	response := utils.BuildSuccessResponse("Get Product Successfully", res)
	ctx.JSON(http.StatusOK, response)
}

