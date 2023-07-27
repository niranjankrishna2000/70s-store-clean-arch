package interfaces

import "main/pkg/utils/models"

type UserRepository interface {
	CheckUserAvailability(email string) bool
	UserBlockStatus(email string) (bool, error)
	FindUserByEmail(user models.UserLogin) (models.UserResponse, error)
	SignUp(user models.UserDetails) (models.UserResponse, error)
}
