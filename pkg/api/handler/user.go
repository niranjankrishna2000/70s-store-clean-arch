package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	services "main/pkg/usecase/interface"
	"main/pkg/utils/models"
	"main/pkg/utils/response"
)

type UserHandler struct {
	userUseCase services.UserUseCase
}

// type Response struct {
// 	ID      uint   `copier:"must"`
// 	Name    string `copier:"must"`
// 	Surname string `copier:"must"`
// }

func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

// @Summary		User Login
// @Description	user can log in by giving their details
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			login  body  models.UserLogin  true	"login"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/login [post]
func (u *UserHandler) Login(c *gin.Context) {
	fmt.Println("=====login handler=====")
	var user models.UserLogin

	if err := c.BindJSON(&user); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	fmt.Println(user)
	user_details, err := u.userUseCase.Login(user)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "User could not be logged in", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	fmt.Println(user_details)


	successRes := response.ClientResponse(http.StatusOK, "User successfully logged in", user_details, nil)
	c.SetCookie("Authorization", user_details.Token, 3600, "", "", true, true)

	c.JSON(http.StatusOK, successRes)
}


// @Summary		User Signup
// @Description	user can signup by giving their details
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			signup  body  models.UserDetails  true	"signup"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/signup [post]
func (u *UserHandler) SignUp(c *gin.Context) {

	var user models.UserDetails
	// bind the user details to the struct
	if err := c.BindJSON(&user); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	// business logic goes inside this function
	userCreated, err := u.userUseCase.SignUp(user)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "User could not signed up", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "User successfully signed up", userCreated, nil)
	c.JSON(http.StatusCreated, successRes)

}