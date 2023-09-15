package dao

import (
	"context"
	"fmt"
	"github.com/cyjhuiz/netflix-clone/backend/notification/model"
	"github.com/cyjhuiz/netflix-clone/backend/notification/netflix_clone_db/public/table"
	"github.com/go-jet/jet/v2/postgres"
)

type NotificationDao struct {
	Store *Store
}

func NewNotificationDao(store *Store) *NotificationDao {
	return &NotificationDao{
		Store: store,
	}
}

func (notificationDao *NotificationDao) GetUserNotificationsByUserID(userID int64) ([]*model.UserNotificationView, error) {
	statement := postgres.
		SELECT(
			table.UserNotification.AllColumns,
			table.Notification.AllColumns,
		).
		FROM(
			table.UserNotification.
				INNER_JOIN(table.Notification, table.UserNotification.NotificationID.EQ(table.Notification.NotificationID)),
		).
		WHERE(
			table.UserNotification.UserID.EQ(postgres.Int(userID)),
		).
		ORDER_BY(
			table.Notification.CreatedAt.DESC(),
		)

	var userNotifications []*model.UserNotificationView

	err := statement.Query(notificationDao.Store.db, &userNotifications)
	if err != nil {
		return nil, err
	}

	return userNotifications, nil
}

func (notificationDao *NotificationDao) CreateAppNotification(userIDs []int64, inputNotification *model.Notification) error {
	ctx := context.Background()
	transaction, err := notificationDao.Store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	fmt.Println(userIDs)

	statement := table.Notification.
		INSERT(table.Notification.MutableColumns.
			Except(table.Notification.CreatedAt)).
		MODEL(inputNotification).
		RETURNING(table.Notification.AllColumns)

	var notifications []*model.Notification
	err = statement.QueryContext(ctx, transaction, &notifications)
	if err != nil {
		return err
	}

	notificationID := notifications[0].NotificationID

	var userNotifications []*model.UserNotification
	for _, userID := range userIDs {
		userNotification := model.NewUserNotification(userID, notificationID)
		userNotifications = append(userNotifications, userNotification)
	}

	statement = table.UserNotification.
		INSERT(table.UserNotification.MutableColumns).
		MODELS(userNotifications)

	printStatementInfo(statement)
	_, err = statement.ExecContext(ctx, transaction)
	if err != nil {
		return err
	}

	transaction.Commit()

	return nil
}

func printStatementInfo(stmt postgres.Statement) {
	query, args := stmt.Sql()

	fmt.Println("Parameterized query: ")
	fmt.Println("==============================")
	fmt.Println(query)
	fmt.Println("Arguments: ")
	fmt.Println(args)

	debugSQL := stmt.DebugSql()

	fmt.Println("\n\nDebug sql: ")
	fmt.Println("==============================")
	fmt.Println(debugSQL)
}
