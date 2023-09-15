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
	err := store.createSubscriptionTable()
	if err != nil {
		return err
	}

	err = store.createUserTable()
	if err != nil {
		return err
	}

	return err
}

func (store *Store) createSubscriptionTable() error {
	query := `CREATE TABLE IF NOT EXISTS subscription(
			subscription_ID SERIAL PRIMARY KEY,
			name VARCHAR(255),
			price FLOAT
		)`
	_, err := store.db.Exec(query)
	return err
}

func (store *Store) createUserTable() error {
	query := `CREATE TABLE IF NOT EXISTS "user"(
			user_ID SERIAL PRIMARY KEY,
			email VARCHAR(255),
			password VARCHAR(100),
			subscription_ID INT,
    		UNIQUE(email),
			CONSTRAINT fk_user_subscription_id FOREIGN KEY (subscription_ID) REFERENCES subscription(subscription_ID)
		)`
	_, err := store.db.Exec(query)
	return err
}
