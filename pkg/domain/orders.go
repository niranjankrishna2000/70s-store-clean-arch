package domain

import "gorm.io/gorm"

type Order struct {
	gorm.Model	`json:"-"`
	UserID        int     `json:"user_id" gorm:"not null"`
	User          User    `json:"-" gorm:"foreignkey:UserID"`
	AddressID     int     `json:"address_id" gorm:"not null"`
	Address       Address `json:"-" gorm:"foreignkey:AddressID"`
	PaymentMethod string  `json:"paymentmethod" gorm:"default:COD"`
	Price         float64 `json:"price"`
	OrderStatus   string  `json:"order_status" gorm:"order_status:4;default:'PENDING';check:order_status IN ('PENDING', 'SHIPPED','DELIVERED','CANCELED')"`
}

type OrderItem struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderID     int       `json:"order_id"`
	Order       Order     `json:"-" gorm:"foreignkey:OrderID;constraint:OnDelete:CASCADE"`
	InventoryID int       `json:"inventory_id"`
	Inventory   Inventory `json:"-" gorm:"foreignkey:InventoryID"`
	Quantity    int       `json:"quantity"`
	TotalPrice  float64   `json:"total_price"`
}

type AdminOrdersResponse struct {
	Pending   []OrderDetails
	Shipped   []OrderDetails
	Delivered []OrderDetails
	Canceled  []OrderDetails
}

type OrderDetails struct {
	Id            int     `json:"order_id" gorm:"column:order_id"`
	Username      string  `json:"name"`
	Address       string  `json:"address"`
	Paymentmethod string  `json:"payment_method" gorm:"column:payment_method"`
	Total         float64 `json:"total"`
}
