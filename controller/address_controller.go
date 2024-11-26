package controller

import (
	"errors"
	"go-foodease-be/dto"
	"go-foodease-be/service"
	"go-foodease-be/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	AddressController interface {
		CreateNewAddress(ctx *gin.Context)
		GetAllAddressByCustomerId(ctx *gin.Context)
		GetAdrressById(ctx *gin.Context)
		UpdateAddressById(ctx *gin.Context)
		DeleteAddressById(ctx *gin.Context)
		GetCustomerActiveAddressById(ctx *gin.Context)
		SetActiveAddress(ctx *gin.Context)
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
		response := utils.BuildFailedResponse("failed process request", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := utils.BuildSuccessResponse("fetch address successfully", addr)
	ctx.JSON(http.StatusOK, response)
}

func (c *addressController) UpdateAddressById(ctx *gin.Context) {
	customerId := ctx.MustGet("id").(string)
	addressId := ctx.Param("address_id")

	if customerId == "" {
		response := utils.BuildFailedResponse("failed to update address", errors.New("cannot identified user"), nil)
		ctx.JSON(http.StatusUnauthorized, response)
		return
	}

	var address dto.CreateNewAddressRequest
	
	if err := ctx.ShouldBind(&address); err != nil {
		response := utils.BuildFailedResponse("failed to update address", err.Error(), nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	res, err := c.addressService.UpdateAddressById(ctx.Request.Context(), addressId, customerId, address)
	if err != nil {
		response := utils.BuildFailedResponse("failed to update address", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildSuccessResponse("success to update address", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *addressController) DeleteAddressById(ctx *gin.Context) {
	customerId := ctx.MustGet("id").(string)
	addressId := ctx.Param("address_id")

	if _, err := c.addressService.GetAddressById(ctx.Request.Context(), addressId, customerId); err != nil {
		response := utils.BuildFailedResponse("failed delete address", err.Error(), nil)
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	if err := c.addressService.DeleteAddressById(ctx.Request.Context(), addressId, customerId); err != nil {
		response := utils.BuildFailedResponse("failed delete address", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildSuccessResponse("success delete address", nil)
	ctx.JSON(http.StatusOK, response)
}

func (c *addressController) GetCustomerActiveAddressById(ctx *gin.Context){
	customerId := ctx.MustGet("id").(string)

	res, err := c.addressService.GetActiveAddress(ctx.Request.Context(), customerId)
	if err != nil {
		response := utils.BuildFailedResponse("failed get active address", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if res.ID == uuid.Nil {
		response := utils.BuildSuccessResponse("there is no active address", struct{}{})
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := utils.BuildSuccessResponse("success get active address", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *addressController) SetActiveAddress(ctx *gin.Context) {
	customerId := ctx.MustGet("id"). (string)
	addressId := ctx.Param("address_id")

	res, err := c.addressService.SetActiveAddress(ctx.Request.Context(), addressId, customerId)
	if err != nil {
		response := utils.BuildFailedResponse("failed to set active address", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildSuccessResponse("success to set active address", res)
	ctx.JSON(http.StatusOK, response)
}