package interfaces

import "main/pkg/utils/models"

type CartUseCase interface {
	AddToCart(user_id, inventory_id int) error
	CheckOut(id int) (models.CheckOut, error)
	ApplyCoupon(userID int,coupon string)(models.CheckOut,error)
}
