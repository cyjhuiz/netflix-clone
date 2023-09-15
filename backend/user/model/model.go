package model

type Subscription struct {
	SubscriptionID int64   `json:"subscriptionID" sql:"primary_key"`
	Name           string  `json:"name"`
	Price          float64 `json:"price"`
}

type User struct {
	UserID         int64  `json:"userID" sql:"primary_key"`
	Email          string `json:"email"`
	Password       string `json:"-"`
	SubscriptionID int64  `json:"subscriptionID"`
}

type UserView struct {
	*User
	Subscription *Subscription `json:"subscription"`
}

func NewUser(email string, password string) *User {
	return &User{
		Email:    email,
		Password: password,
	}
}
