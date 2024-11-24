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
		Register(ctx *gin.Context)
		DeleteAccount(ctx *gin.Context)
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
	customerId := ctx.MustGet("id").(string)	
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

func (c *customerController) Register(ctx *gin.Context) {
	var customer dto.CustomerRegisterRequest
	if err := ctx.ShouldBind(&customer); err != nil {
		response := utils.BuildFailedResponse("Failed Get Data From Body", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	res, err := c.customerService.Register(ctx.Request.Context(), customer)
	fmt.Println("err ctrl", err)
	if err != nil {
		response := utils.BuildFailedResponse("Failed Register Customer", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return 
	}

	response := utils.BuildSuccessResponse("Customer Registered Successfully", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *customerController) DeleteAccount(ctx *gin.Context) {
	customerId := ctx.MustGet("id").(string)
	err := c.customerService.DeleteAccount(ctx.Request.Context(), customerId)
	if err != nil {
		response := utils.BuildFailedResponse("failed to delete account", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildSuccessResponse("success to delete account", nil)
	ctx.JSON(http.StatusOK, response)
}