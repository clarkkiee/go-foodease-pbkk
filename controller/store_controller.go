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
	StoreController interface {
		Login(ctx *gin.Context)
		Register(ctx *gin.Context)
		DeleteAccount(ctx *gin.Context)
	}

	storeController struct {
		storeService service.StoreService
		addressService service.AddressService
	}
)

func NewStoreController(ss service.StoreService, as service.AddressService) StoreController {
	return &storeController{
		storeService: ss,
		addressService: as,
	}
}

func (c *storeController) Login(ctx *gin.Context){
	var req dto.StoreLoginRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response := utils.BuildFailedResponse("invalid login request schema", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

		res, err := c.storeService.VerifyLogin(ctx.Request.Context(), req)
		if err != nil {
			response := utils.BuildFailedResponse("failed login", err.Error(), nil)
			ctx.JSON(http.StatusBadRequest, response)
			return
		}

		response := utils.BuildSuccessResponse("login success", res)
		ctx.JSON(http.StatusOK, response)
}

func (c *storeController) Register(ctx *gin.Context) {
	var store dto.StoreRegisterRequest
	var address dto.CreateNewAddressRequest

	if err := ctx.ShouldBind(&store); err != nil {
		response := utils.BuildFailedResponse("Failed Get Data From Body", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	resAddress, errAddress := c.addressService.CreateNewStoreAddress(ctx, address)
	fmt.Println(resAddress)
	if errAddress != nil {
		response := utils.BuildFailedResponse("Failed Register Store Address", errAddress.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	res, err := c.storeService.RegisterAccount(ctx, store, resAddress)
	if err != nil {
		response := utils.BuildFailedResponse("Failed Register Store", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildSuccessResponse("Store Registered Successfully", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *storeController) DeleteAccount(ctx *gin.Context) {
	storeId := ctx.MustGet("id").(string)
	err := c.storeService.DeleteAccount(ctx.Request.Context(), storeId)
	if err != nil {
		response := utils.BuildFailedResponse("failed to delete account", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildSuccessResponse("success to delete account", nil)
	ctx.JSON(http.StatusOK, response)
}

