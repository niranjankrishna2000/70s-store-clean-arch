package interfaces

import (
	"main/pkg/domain"
)

type CartRepository interface {
	//GetCart(id int) ([]models.GetCart, error)
	GetAddresses(id int) ([]domain.Address, error)
	CheckIfInvAdded(invID,cartID int)bool
	GetCartId(user_id int) (int, error)
	CreateNewCart(user_id int) (int, error)
	AddLineItems(invID,cartID int) error
	AddQuantity(invID,cartID int) error
}
