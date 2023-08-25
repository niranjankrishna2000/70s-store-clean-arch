package repository

import (
	"errors"
	"main/pkg/domain"
	interfaces "main/pkg/repository/interface"

	"gorm.io/gorm"
)

type cartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) interfaces.CartRepository {
	return &cartRepository{
		DB: db,
	}
}

func (ad *cartRepository) GetAddresses(id int) ([]domain.Address, error) {

	var addresses []domain.Address

	if err := ad.DB.Raw("SELECT * FROM addresses WHERE user_id=?", id).Scan(&addresses).Error; err != nil {
		return []domain.Address{}, err
	}

	return addresses, nil

}

func (ad *cartRepository) GetCartId(user_id int) (int, error) {

	var id int

	if err := ad.DB.Raw("SELECT id FROM carts WHERE user_id=?", user_id).Scan(&id).Error; err != nil {
		return 0, err
	}

	return id, nil

}

func (i *cartRepository) CreateNewCart(user_id int) (int, error) {
	var id int
	err := i.DB.Exec(`
		INSERT INTO carts (user_id)
		VALUES (?)`, user_id).Error
	if err != nil {
		return 0, err
	}

	if err := i.DB.Raw("select id from carts where user_id=?", user_id).Scan(&id).Error; err != nil {
		return 0, err
	}

	return id, nil
}

func (i *cartRepository) AddLineItems(cart_id, inventory_id int) error {

	err := i.DB.Exec(`
		INSERT INTO line_items (cart_id,inventory_id)
		VALUES (?,?)`, cart_id, inventory_id).Error
	if err != nil {
		return err
	}

	return nil
}

func (i *cartRepository) ValidateCoupon(coupon string) (bool, error){
	count:=0
	if err := i.DB.Raw("select count(id) from coupons where name=?", coupon).Scan(&count).Error; err != nil {
		return false, err
	}
	if count<1{
		return false,errors.New("not a valid coupon")
	}
	valid:=true
	if err := i.DB.Raw("select valid from coupons where name=?", coupon).Scan(&valid).Error; err != nil {
		return false, err
	}
	if !valid{
		return false,errors.New("not a valid coupon")
	}
	return true,nil
}


func (i *cartRepository) GetDiscountRate(coupon string) (int, error){
	discount:=0
	if err := i.DB.Raw("select discount_rate from coupons where name=?", coupon).Scan(&discount).Error; err != nil {
		return 0, err
	}
	return discount,nil
}
