package model

type CreateUserRequest struct {
	Email    string
	Password string
}

type LoginUserRequest struct {
	Email    string
	Password string
}

type LoginUserResponse struct {
	UserId int64  `json:"userID"`
	Token  string `json:"token"`
}
