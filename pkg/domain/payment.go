package domain

type PaymentMethod struct {
	ID            int    `gorm:"primaryKey"`
	PaymentMethod string `json:"PaymentMethod" validate:"required" gorm:"unique"`
}
