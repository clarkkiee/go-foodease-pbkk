package controller

import (
	"go-foodease-be/dto"
	"go-foodease-be/service"
	"go-foodease-be/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	OrderController interface {
		AddToCart(ctx *gin.Context)
	}

	orderController struct {
		orderService service.OrderService
		productService service.ProductService
	}
)

func NewOrderController(os service.OrderService, ps service.ProductService) OrderController {
	return &orderController{
		orderService: os,
		productService: ps,
	}
}

func (c *orderController) AddToCart(ctx *gin.Context) {
	customerId := ctx.MustGet("id").(string)

	var productReq dto.AddToCartSchema
	if err := ctx.ShouldBind(&productReq); err != nil {
		response := utils.BuildFailedResponse("failed to add product to cart", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productInfo, err := c.productService.GetMinimumProduct(ctx.Request.Context(), productReq.ProductId)
	if err != nil {
		response := utils.BuildFailedResponse("failed to add product to cart", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	c.orderService.AddToCart(ctx.Request.Context(), productInfo.StoreID, customerId)	
}