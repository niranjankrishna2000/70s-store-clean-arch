package interfaces

import (
	"main/pkg/domain"
	"main/pkg/utils/models"
)

type UserRepository interface {
	CheckUserAvailability(email string) bool
	UserBlockStatus(email string) (bool, error)
	FindUserByEmail(user models.UserLogin) (models.UserResponse, error)
	SignUp(user models.UserDetails) (models.UserResponse, error)
	AddAddress(id int, address models.AddAddress, result bool) error
	GetAddresses(id int) ([]domain.Address, error)
	CheckIfFirstAddress(id int) bool
	GetUserDetails(id int) (models.UserResponse, error)
	ChangePassword(id int, password string) error
	GetPassword(id int) (string, error)
	FindIdFromPhone(phone string) (int, error)
	EditName(id int, name string) error
	EditEmail(id int, email string) error
	EditPhone(id int, phone string) error
}
