package interfaces

import (
	"main/pkg/domain"
)

type CouponRepository interface {
	AddCoupon(domain.Coupon) error
	MakeCouponInvalid(id int) error
	FindCouponDiscount(couponID int) int
	GetCoupons(page,limit int) ([]domain.Coupon,error)
}
