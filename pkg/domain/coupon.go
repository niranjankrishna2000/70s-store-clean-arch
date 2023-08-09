package domain

type Coupon struct {
	ID           int `json:"-" gorm:"primarykey"`
	Name         int  `json:"name" gorm:"unique;not null"`
	DiscountRate int  `json:"discount_rate"`
	Valid        bool `gorm:"default:True"`
}
