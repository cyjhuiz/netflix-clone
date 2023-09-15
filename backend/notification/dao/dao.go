package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Store struct {
	db *sql.DB
}

func NewStore() (*Store, error) {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "password"
		dbname   = "netflix_clone_db"
	)

	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: db,
	}, nil
}

func (store *Store) Init() error {
	err := store.createNotificationTable()
	if err != nil {
		return err
	}

	err = store.createUserNotificationTable()
	if err != nil {
		return err
	}

	return nil
}

func (store *Store) createNotificationTable() error {
	query := `CREATE TABLE IF NOT EXISTS notification(
		notification_id SERIAL PRIMARY KEY,
		title VARCHAR(100),
    	description VARCHAR(100),
		thumbnail_url VARCHAR(255),
    	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`

	_, err := store.db.Exec(query)
	return err
}

func (store *Store) createUserNotificationTable() error {
	query := `CREATE TABLE IF NOT EXISTS user_notification(
    	user_notification_id SERIAL PRIMARY KEY,
		user_id INTEGER,
        notification_id INTEGER,
        UNIQUE(user_id, notification_id),
        CONSTRAINT fk_notification_notification_id FOREIGN KEY (notification_id) REFERENCES notification(notification_id)
	)`

	_, err := store.db.Exec(query)
	return err
}
