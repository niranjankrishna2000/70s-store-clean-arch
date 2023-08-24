package usecase

import (
	"main/pkg/domain"
	interfaces "main/pkg/repository/interface"
	services "main/pkg/usecase/interface"
)

type couponUseCase struct {
	couponRepo interfaces.CouponRepository
}

func NewCouponUseCase(couponRepo interfaces.CouponRepository) services.CouponUseCase{
	return &couponUseCase{
		couponRepo: couponRepo,
	}
}

func (coup *couponUseCase) AddCoupon(coupon domain.Coupon) error {
	if err := coup.couponRepo.AddCoupon(coupon); err != nil {
		return err
	}

	return nil
}

func (coup *couponUseCase) MakeCouponInvalid(id int) error {
	if err := coup.couponRepo.MakeCouponInvalid(id); err != nil {
		return err
	}

	return nil
}
func (coup *couponUseCase) GetCoupons(page,limit int)([]domain.Coupon,error){
	coupons,err:=coup.couponRepo.GetCoupons(page,limit)
	if err != nil {
		return []domain.Coupon{}, err
	}
	return coupons,nil
}