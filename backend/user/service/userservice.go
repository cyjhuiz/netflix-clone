package service

import (
	"context"
	"fmt"
	"github.com/cyjhuiz/netflix-clone/backend/user/dao"
	"github.com/cyjhuiz/netflix-clone/backend/user/model"
	"github.com/cyjhuiz/netflix-clone/backend/user/util/authutil"
)

type UserService struct {
	UserDao *dao.UserDao
}

func NewUserService(userDao *dao.UserDao) *UserService {
	return &UserService{
		UserDao: userDao,
	}
}

func (userService *UserService) GetUsers(ctx context.Context) ([]*model.UserView, error) {
	userView, err := userService.UserDao.GetUsers()
	if err != nil {
		return nil, err
	}

	return userView, err
}

func (userService *UserService) GetUsersByEmail(ctx context.Context, email string) ([]*model.UserView, error) {
	userView, err := userService.UserDao.GetUsersByEmail(email)
	if err != nil {
		return nil, err
	}

	return userView, err
}

func (userService *UserService) GetUsersByUserIDs(ctx context.Context, userIDs []int64) ([]*model.UserView, error) {
	user, err := userService.UserDao.GetUserViewsByUserIDs(userIDs)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (userService *UserService) GetUserByUserID(ctx context.Context, userID int64) (*model.UserView, error) {
	userView, err := userService.UserDao.GetUserViewByUserID(userID)
	if err != nil {
		return nil, err
	}

	return userView, err
}

func (userService *UserService) CreateUser(ctx context.Context, user *model.User) error {
	users, err := userService.UserDao.GetUsersByEmail(user.Email)
	if len(users) > 0 {
		return fmt.Errorf("email already exists")
	} else if err != nil {
		return err
	}

	encryptedPassword, err := authutil.EncryptPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = encryptedPassword

	err = userService.UserDao.CreateUser(user)
	if err != nil {
		return err
	}

	return err
}

func (userService *UserService) LoginUser(ctx context.Context, email string, password string) (*model.LoginUserResponse, error) {
	users, err := userService.UserDao.GetUsersByEmail(email)
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, fmt.Errorf("user not found")
	}

	user := users[0]

	err = authutil.ValidatePassword(password, user.Password)
	if err != nil {
		return nil, fmt.Errorf("invalid authentication details")
	}

	token, err := authutil.CreateJWT(user.UserID)
	if err != nil {
		return nil, err
	}

	loginUserResponse := &model.LoginUserResponse{
		UserId: user.UserID,
		Token:  token,
	}

	return loginUserResponse, nil
}
