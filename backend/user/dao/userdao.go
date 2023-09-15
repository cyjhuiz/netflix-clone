package dao

import (
	"fmt"
	"github.com/cyjhuiz/netflix-clone/backend/user/model"
	"github.com/cyjhuiz/netflix-clone/backend/user/netflix_clone_db/public/table"
	"github.com/go-jet/jet/v2/postgres"
	_ "github.com/lib/pq"
)

// Get User By ID
// Create User
// Login User

type UserDao struct {
	Store *Store
}

func NewUserDao(store *Store) *UserDao {
	return &UserDao{
		Store: store,
	}
}

func (userDao *UserDao) GetUsers() ([]*model.UserView, error) {
	statement := postgres.
		SELECT(
			table.User.AllColumns,
			table.Subscription.AllColumns,
		).
		FROM(
			table.User.
				LEFT_JOIN(table.Subscription, table.User.SubscriptionID.EQ(table.Subscription.SubscriptionID)),
		)

	var users []*model.UserView
	err := statement.Query(userDao.Store.db, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (userDao *UserDao) GetUsersByEmail(email string) ([]*model.UserView, error) {
	statement := postgres.
		SELECT(
			table.User.AllColumns,
			table.Subscription.AllColumns,
		).
		FROM(
			table.User.
				LEFT_JOIN(table.Subscription, table.User.SubscriptionID.EQ(table.Subscription.SubscriptionID)),
		).WHERE(
		table.User.Email.EQ(postgres.String(email)),
	)

	var users []*model.UserView
	err := statement.Query(userDao.Store.db, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (userDao *UserDao) GetUserViewsByUserIDs(userIDs []int64) ([]*model.UserView, error) {
	var userIDsInPostgresFormat []postgres.Expression
	for _, userID := range userIDs {
		userIDsInPostgresFormat = append(userIDsInPostgresFormat, postgres.Int(userID))
	}

	statement := postgres.
		SELECT(
			table.User.AllColumns,
			table.Subscription.AllColumns,
		).
		FROM(
			table.User.
				LEFT_JOIN(table.Subscription, table.User.SubscriptionID.EQ(table.Subscription.SubscriptionID)),
		).WHERE(
		table.User.UserID.IN(userIDsInPostgresFormat...),
	)

	var users []*model.UserView
	err := statement.Query(userDao.Store.db, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (userDao *UserDao) GetUserViewByUserID(userID int64) (*model.UserView, error) {
	statement := postgres.
		SELECT(
			table.User.AllColumns,
			table.Subscription.AllColumns,
		).
		FROM(
			table.User.
				LEFT_JOIN(table.Subscription, table.User.SubscriptionID.EQ(table.Subscription.SubscriptionID)),
		).WHERE(
		table.User.UserID.EQ(postgres.Int(userID)),
	)

	var users []*model.UserView
	err := statement.Query(userDao.Store.db, &users)
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return users[0], nil
}

func (userDao *UserDao) CreateUser(user *model.User) error {
	statement := table.User.
		INSERT(
			table.User.MutableColumns.
				Except(table.User.SubscriptionID)).
		MODEL(user)

	_, err := statement.Exec(userDao.Store.db)
	return err
}
