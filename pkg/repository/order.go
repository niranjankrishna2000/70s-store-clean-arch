package repository

import (
	"errors"
	"main/pkg/domain"
	interfaces "main/pkg/repository/interface"
	"main/pkg/utils/models"

	"gorm.io/gorm"
)

type orderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) interfaces.OrderRepository {
	return &orderRepository{
		DB: db,
	}
}

func (or *orderRepository) GetOrders(id int) ([]domain.Order, error) {

	var orders []domain.Order

	if err := or.DB.Raw("select * from orders where user_id=?", id).Scan(&orders).Error; err != nil {
		return []domain.Order{}, err
	}

	return orders, nil

}

func (ad *orderRepository) GetCart(id int) ([]models.GetCart, error) {

	var cart []models.GetCart

	if err := ad.DB.Raw("SELECT inventories.product_name,cart_products.quantity,cart_products.total_price AS Total FROM cart_products JOIN inventories ON cart_products.inventory_id=inventories.id WHERE user_id=?", id).Scan(&cart).Error; err != nil {
		return []models.GetCart{}, err
	}
	return cart, nil

}

func (i *orderRepository) OrderItems(userid int, addressid int, total float64) (int, error) {

	var id int
	query := `
    INSERT INTO orders (user_id,address_id,  price)
    VALUES (?, ?, ?)
    RETURNING id
    `
	i.DB.Raw(query, userid, addressid, total).Scan(&id)

	return id, nil

}

func (i *orderRepository) AddOrderProducts(order_id int, cart []models.GetCart) error {

	query := `
    INSERT INTO order_items (order_id,inventory_id,quantity,total_price)
    VALUES (?, ?, ?, ?)
    `

	for _, v := range cart {
		var inv int
		if err := i.DB.Raw("select id from inventories where product_name=?", v.ProductName).Scan(&inv).Error; err != nil {
			return err
		}

		if err := i.DB.Exec(query, order_id, inv, v.Quantity, v.Total).Error; err != nil {
			return err
		}
	}

	return nil

}

func (i *orderRepository) CancelOrder(id int) error {

	if err := i.DB.Exec("update orders set order_status='CANCELED' where id=?", id).Error; err != nil {
		return err
	}

	return nil

}

func (i *orderRepository) EditOrderStatus(status string, id int) error {

	if err := i.DB.Exec("update orders set order_status=? where id=?", status, id).Error; err != nil {
		return err
	}

	return nil

}

func (or *orderRepository) AdminOrders(status string) ([]domain.OrderDetails, error) {

	var orders []domain.OrderDetails
	if err := or.DB.Raw("SELECT orders.id AS order_id, users.name AS username, CONCAT(addresses.house_name, ' ', addresses.street, ' ', addresses.city) AS address, orders.payment_method AS payment_method, orders.price As total FROM orders JOIN users ON users.id = orders.user_id JOIN addresses ON orders.address_id = addresses.id WHERE order_status = ?", status).Scan(&orders).Error; err != nil {
		return []domain.OrderDetails{}, err
	}

	return orders, nil

}

func (o *orderRepository) CheckOrder(orderID string, userID int) error {

	var count int
	err := o.DB.Raw("select count(*) from orders where order_id = ?", orderID).Scan(&count).Error
	if err != nil {
		return err
	}
	if count < 0 {
		return errors.New("no such order exist")
	}
	var checkUser int
	err = o.DB.Raw("select user_id from orders where order_id = ?", orderID).Scan(&checkUser).Error
	if err != nil {
		return err
	}

	if userID != checkUser {
		return errors.New("the order is not did by this user")
	}

	return nil
}

func (o *orderRepository) GetOrderDetail(orderID string) (domain.Order, error) {

	var orderDetails domain.Order
	err := o.DB.Raw("select * from orders where order_id = ?", orderID).Scan(&orderDetails).Error
	if err != nil {
		return domain.Order{}, err
	}

	return orderDetails, nil

}

func (o *orderRepository) FindAmountFromOrderID(id int) (float64, error) {

	var amount float64
	err := o.DB.Raw("select final_price from orders where id = ?", id).Scan(&amount).Error
	if err != nil {
		return 0, err
	}

	return amount, nil
}

func (o *orderRepository) FindUserIdFromOrderID(id int) (int, error) {

	var user_id int
	err := o.DB.Raw("select user_id from orders where id = ?", id).Scan(&user_id).Error
	if err != nil {
		return 0, err
	}

	return user_id, nil
}
