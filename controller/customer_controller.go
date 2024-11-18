package controller

import (
	"fmt"
	"go-foodease-be/dto"
	"go-foodease-be/service"
	"go-foodease-be/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	CustomerController interface {
		Me(ctx *gin.Context)
		Login(ctx *gin.Context)
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

func (c *customerController) Login(ctx *gin.Context){
	var req dto.CustomerLoginRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response := utils.BuildFailedResponse("invalid login request schema", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

		res, err := c.customerService.VerifyLogin(ctx.Request.Context(), req)
		if err != nil {
			response := utils.BuildFailedResponse("failed login", err.Error(), nil)
			ctx.JSON(http.StatusBadRequest, response)
			return
		}

		response := utils.BuildSuccessResponse("login success", res)
		ctx.JSON(http.StatusOK, response)
}