package handler

import (
	"main/pkg/domain"
	services "main/pkg/usecase/interface"
	"main/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CouponHandler struct {
	couponUsecase services.CouponUseCase
}

func NewCouponHandler(coupUsecase services.CouponUseCase) *CouponHandler {
	return &CouponHandler{
		couponUsecase: coupUsecase,
	}
}

// @Summary		Add Coupon
// @Description	Admin can add new coupons
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			coupon	body	domain.Coupon	true	"coupon"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/coupons/create [post]
func (coup *CouponHandler) CreateNewCoupon(c *gin.Context) {
	var coupon domain.Coupon
	if err := c.BindJSON(&coupon); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	err := coup.couponUsecase.AddCoupon(coupon)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not add the Coupon", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added Coupon", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary		Make Coupon invalid
// @Description	Admin can make the coupons as invalid so that users cannot use that particular coupon
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			id	query	string	true	"id"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/coupons/expire [post]
func (coup *CouponHandler) MakeCouponInvalid(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := coup.couponUsecase.MakeCouponInvalid(id); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Coupon cannot be made invalid", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully made Coupon as invalid", nil, nil)
	c.JSON(http.StatusOK, successRes)

}
