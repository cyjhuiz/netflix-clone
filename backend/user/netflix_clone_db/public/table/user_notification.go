//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var UserNotification = newUserNotificationTable("public", "user_notification", "")

type userNotificationTable struct {
	postgres.Table

	// Columns
	UserID         postgres.ColumnInteger
	NotificationID postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type UserNotificationTable struct {
	userNotificationTable

	EXCLUDED userNotificationTable
}

// AS creates new UserNotificationTable with assigned alias
func (a UserNotificationTable) AS(alias string) *UserNotificationTable {
	return newUserNotificationTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new UserNotificationTable with assigned schema name
func (a UserNotificationTable) FromSchema(schemaName string) *UserNotificationTable {
	return newUserNotificationTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new UserNotificationTable with assigned table prefix
func (a UserNotificationTable) WithPrefix(prefix string) *UserNotificationTable {
	return newUserNotificationTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new UserNotificationTable with assigned table suffix
func (a UserNotificationTable) WithSuffix(suffix string) *UserNotificationTable {
	return newUserNotificationTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newUserNotificationTable(schemaName, tableName, alias string) *UserNotificationTable {
	return &UserNotificationTable{
		userNotificationTable: newUserNotificationTableImpl(schemaName, tableName, alias),
		EXCLUDED:              newUserNotificationTableImpl("", "excluded", ""),
	}
}

func newUserNotificationTableImpl(schemaName, tableName, alias string) userNotificationTable {
	var (
		UserIDColumn         = postgres.IntegerColumn("user_id")
		NotificationIDColumn = postgres.IntegerColumn("notification_id")
		allColumns           = postgres.ColumnList{UserIDColumn, NotificationIDColumn}
		mutableColumns       = postgres.ColumnList{UserIDColumn, NotificationIDColumn}
	)

	return userNotificationTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		UserID:         UserIDColumn,
		NotificationID: NotificationIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
