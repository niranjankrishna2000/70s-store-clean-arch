package usecase

import (
	"errors"
	"fmt"
	"main/pkg/helper"
	interfaces "main/pkg/repository/interface"
	services "main/pkg/usecase/interface"
	"main/pkg/utils/models"

	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

func (u *userUseCase) Login(user models.UserLogin) (models.TokenUser, error) {
	fmt.Println("=====login usecase=====")
	// checking if a username exist with this email address
	ok := u.userRepo.CheckUserAvailability(user.Email)
	if !ok {
		return models.TokenUser{}, errors.New("the user does not exist")
	}

	isBlocked, err := u.userRepo.UserBlockStatus(user.Email)
	if err != nil {
		return models.TokenUser{}, err
	}

	if isBlocked {
		return models.TokenUser{}, errors.New("user is blocked by admin")
	}

	// Get the user details in order to check the password, in this case ( The same function can be reused in future )
	user_details, err := u.userRepo.FindUserByEmail(user)
	if err != nil {
		return models.TokenUser{}, err
	}
	fmt.Println(user_details)
	err = bcrypt.CompareHashAndPassword([]byte(user_details.Password), []byte(user.Password))
	if err != nil {
		return models.TokenUser{}, errors.New("password incorrect")
	}
	
	tokenString, err := helper.GenerateTokenUser(user_details)
	if err != nil {
		return models.TokenUser{}, errors.New("could not create token")
	}
	fmt.Println(tokenString)
	return models.TokenUser{
		Username: user_details.Username,
		Token:    tokenString,
	}, nil

}

func (u *userUseCase) SignUp(user models.UserDetails) (models.TokenUser, error) {
	// Check whether the user already exist. If yes, show the error message, since this is signUp
	userExist := u.userRepo.CheckUserAvailability(user.Email)
	if userExist {
		return models.TokenUser{}, errors.New("user already exist, sign in")
	}
	if user.Password != user.ConfirmPassword {
		return models.TokenUser{}, errors.New("password does not match")
	}

	// Hash password since details are validated

	hashedPassword, err := helper.PasswordHashing(user.Password)
	if err != nil {
		return models.TokenUser{}, err
	}

	user.Password = hashedPassword

	// add user details to the database
	userData, err := u.userRepo.SignUp(user)
	if err != nil {
		return models.TokenUser{}, err
	}

	// crete a JWT token string for the user
	tokenString, err := helper.GenerateTokenUser(userData)
	if err != nil {
		return models.TokenUser{}, errors.New("could not create token due to some internal error")
	}

	
	return models.TokenUser{
		Username: user.Username,
		Token: tokenString,
	}, nil
}
