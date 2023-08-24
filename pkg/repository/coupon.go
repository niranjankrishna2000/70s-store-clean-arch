package repository

import (
	"main/pkg/domain"
	interfaces "main/pkg/repository/interface"

	"gorm.io/gorm"
)

type couponRepository struct {
	db *gorm.DB
}

func NewCouponRepository(DB *gorm.DB) interfaces.CouponRepository {
	return &couponRepository{
		db: DB,
	}
}

func (c *couponRepository) AddCoupon(coup domain.Coupon) error {
	if err := c.db.Exec("INSERT INTO coupons(name,discount_rate,valid) values($1,$2,$3)", coup.Name, coup.DiscountRate, coup.Valid).Error; err != nil {
		return err
	}

	return nil
}

func (c *couponRepository) MakeCouponInvalid(id int) error {
	if err := c.db.Exec("UPDATE coupons SET valid=false where id=$1", id).Error; err != nil {
		return err
	}

	return nil
}

func (c *couponRepository) FindCouponDiscount(couponID int) int {
	var coupon domain.Coupon
	err := c.db.Raw("select name,discount_rate,valid from coupons where id=$1", couponID).Scan(&coupon).Error
	if err != nil {
		return 0
	}
	if !coupon.Valid{
		return 1
	}

	return coupon.DiscountRate
}

func (c *couponRepository) GetCoupons(page,limit int) ([]domain.Coupon, error){
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	offset := (page - 1) * limit
	var coupons []domain.Coupon

	if err := c.db.Raw("select id,name,discount_rate,valid from coupons limit ? offset ?", limit, offset).Scan(&coupons).Error; err != nil {
		return []domain.Coupon{}, err
	}

	return coupons, nil
}