package response

import "reglog/internal/model"

type User struct {
	ID          uint   `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	Username    string `json:"username" form:"username"`
	Email       string `json:"email" from:"email"`
	Gender      string `json:"gender" form:"gender"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Picture     string `json:"picture" form:"picture"`
}

type LoginUser struct {
	ID       uint   `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Token    string `json:"token" from:"token"`
}

type GetUser struct {
	ID       uint   `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Username string `json:"username" form:"username"`
}

func FromModel(model model.User) User {
	return User{
		ID:       model.ID,
		Name:     model.Name,
		Username: model.Username,
		Email:    model.Email,
	}
}
