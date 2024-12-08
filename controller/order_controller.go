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
		GetCustomerCart(ctx *gin.Context)
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

	res, err := c.orderService.AddToCart(ctx.Request.Context(), customerId, productReq.ProductId)	
	if err != nil {
		response := utils.BuildFailedResponse("failed to add product to cart", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildSuccessResponse("success to add product to cart", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *orderController) GetCustomerCart(ctx *gin.Context) {
	customerId := ctx.MustGet("id").(string)
	res, err := c.orderService.GetCustomerCart(ctx.Request.Context(), customerId)
	if err != nil {
		response := utils.BuildFailedResponse("failed to get customer cart", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}	

	response := utils.BuildSuccessResponse("success to get customer cart", res)
	ctx.JSON(http.StatusOK, response)
}