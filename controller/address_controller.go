package controller

import (
	"errors"
	"go-foodease-be/dto"
	"go-foodease-be/service"
	"go-foodease-be/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	AddressController interface {
		CreateNewAddress(ctx *gin.Context)
		GetAllAddressByCustomerId(ctx *gin.Context)
		GetAdrressById(ctx *gin.Context)
	}

	addressController struct {
		addressService service.AddressService
	}
)

func NewAddressController(as service.AddressService) AddressController {
	return &addressController{
		addressService: as,
	}
}

func (c *addressController) CreateNewAddress(ctx *gin.Context){

	customerId := ctx.MustGet("id").(string)
	if customerId == "" {
		response := utils.BuildFailedResponse("failed to create address", errors.New("user not authorized"), nil)
		ctx.JSON(http.StatusUnauthorized, response)
		return
	}

	var req dto.CreateNewAddressRequest

	if err := ctx.ShouldBind(&req); err != nil {
		response := utils.BuildFailedResponse("invalid create address request schema", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	addr, err := c.addressService.CreateNewAddress(ctx.Request.Context(), req, customerId)
	if err != nil {
		response := utils.BuildFailedResponse("failed to create address", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildSuccessResponse("success creating new address", addr)
	ctx.JSON(http.StatusOK, response)
}

func (c *addressController) GetAllAddressByCustomerId(ctx *gin.Context){
	customerId := ctx.MustGet("id").(string)
	if customerId == "" {
		response := utils.BuildFailedResponse("fetch address failed", errors.New("cannot identified user"), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	addresses, err := c.addressService.GetAllAddressByCustomerId(ctx.Request.Context(), customerId)
	if err != nil {
		response := utils.BuildFailedResponse("fetch address failed", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := utils.BuildSuccessResponse("fetch address successful", addresses)
	ctx.JSON(http.StatusOK, response)
}

func (c *addressController) GetAdrressById(ctx *gin.Context){
	customerId := ctx.MustGet("id").(string)
	addressId := ctx.Param("address_id")
	
	if addressId == "" {
		response := utils.BuildFailedResponse("failed process request", errors.New("must specify address_id"), nil)
		ctx.JSON(http.StatusBadGateway, response)
		return
	}

	addr, err := c.addressService.GetAddressById(ctx.Request.Context(), addressId, customerId)
	if err != nil {
		response := utils.BuildFailedResponse("failed process request", errors.New("failed to get address"), nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := utils.BuildSuccessResponse("fetch address successfully", addr)
	ctx.JSON(http.StatusOK, response)
}