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
