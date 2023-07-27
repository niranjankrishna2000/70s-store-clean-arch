package domain

type User struct {
	Id         int    `gorm:"primaryKey"`
	Name       string `json:"name"`
	Email      string `gorm:"unique" json:"email"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Phone      string `gorm:"unique" json:"phone"`
	Permission bool   `gorm:"default:true" json:"permission"`
}

type Address struct {
	Id        uint   `json:"id" gorm:"unique;not null"`
	UserID    uint   `json:"user_id"`
	user	  User  `json:"-" gorm:"foreignkey:UserID"`
	Name      string `json:"name" validate:"required"`
	HouseName string `json:"house_name" validate:"required"`
	Street    string `json:"street" validate:"required"`
	City      string `json:"city" validate:"required"`
	State     string `json:"state" validate:"required"`
	Pin       string `json:"pin" validate:"required"`
	Default   bool   `json:"default" gorm:"default:false"`
}