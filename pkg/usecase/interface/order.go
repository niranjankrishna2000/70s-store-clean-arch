package interfaces

import (
	"main/pkg/domain"
)

type OrderUseCase interface {
	GetOrders(id int) ([]domain.Order, error)
	OrderItemsFromCart(userid int, addressid int) error
	CancelOrder(id int) error
	EditOrderStatus(status string, id int) error
	AdminOrders() (domain.AdminOrdersResponse, error)
}
