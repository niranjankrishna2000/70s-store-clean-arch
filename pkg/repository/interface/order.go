package interfaces

import (
	"main/pkg/domain"
	"main/pkg/utils/models"
	"time"
)

type OrderRepository interface {
	GetOrders(id, page, limit int) ([]domain.Order, error)
	GetProductsQuantity() ([]domain.ProductReport, error)
	GetOrdersInRange(startDate, endDate time.Time) ([]domain.Order, error)
	GetProductNameFromID(id int) (string, error)
	GetCart(userid int) ([]models.GetCart, error)

	OrderItems(userid int, order models.Order, total float64) (int, error)
	AddOrderProducts(order_id int, cart []models.GetCart) error
	CancelOrder(orderid int) error
	EditOrderStatus(status string, id int) error
	MarkAsPaid(orderID int) error
	AdminOrders(page, limit int, status string) ([]domain.OrderDetails, error)

	CheckOrder(orderID string, userID int) error
	GetOrderDetail(orderID string) (domain.Order, error)
	FindUserIdFromOrderID(orderID int) (int, error)
	FindAmountFromOrderID(orderID int) (float64, error)
	ReturnOrder(id int) error
	CheckOrderStatus(orderID int) (string, error)
	CheckPaymentStatus(orderID int) (string, error)
}
