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