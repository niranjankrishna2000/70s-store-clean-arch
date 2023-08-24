package interfaces

import "main/pkg/domain"

type CouponUseCase interface {
	AddCoupon(coupon domain.Coupon) error
	MakeCouponInvalid(id int) error
	GetCoupons(page,limit int)([]domain.Coupon,error)

}
