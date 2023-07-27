package models


type UserLogin struct{
	Email 		string `json:"email"`
	Password	string `json:"password"`
}

type TokenUser struct{
	Username 	string
	Token		string
}

type UserResponse struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email" validate:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

//signup
type UserDetails struct {
	Name            string `json:"name"`
	Email           string `json:"email" validate:"email"`
	Username 		string `json:"username"`
	Phone           string `json:"phone"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmpassword"`
}