package controller

import (
	"fmt"
	"go-foodease-be/service"

	"github.com/gin-gonic/gin"
)

type (
	CustomerController interface {
		Me(ctx *gin.Context)
	}

	customerController struct {
		customerService service.CustomerService
	}
)

func NewCustomerController(cs service.CustomerService) CustomerController {
	return &customerController{
		customerService: cs,
	}
}

func (c *customerController) Me(ctx *gin.Context) {
	// customerId := ctx.MustGet("customer_id").(string)
	customerId := "f577a67f-e6eb-46c0-866e-3fde4eaa185a"
	customer, err := c.customerService.GetCustomerById(ctx.Request.Context(), customerId)
	if err != nil {
		fmt.Println(err)
	}

	ctx.JSON(200, customer)
}