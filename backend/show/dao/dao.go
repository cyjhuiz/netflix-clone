package dao

import (
	"database/sql"
	"fmt"
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
	err := store.createCategoryTable()
	if err != nil {
		return err
	}

	err = store.createShowTable()
	if err != nil {
		return err
	}

	err = store.createEpisodeTable()
	if err != nil {
		return err
	}

	err = store.createLikeTable()
	if err != nil {
		return err
	}

	err = store.createFavouriteTable()
	if err != nil {
		return err
	}

	return err
}

// database table creation - from tables with no dependencies to tables with dependencies
func (store *Store) createCategoryTable() error {
	query := `CREATE TABLE IF NOT EXISTS category(
		category_id SERIAL PRIMARY KEY,
		name VARCHAR(100)
	)`

	_, err := store.db.Exec(query)
	return err
}

func (store *Store) createShowTable() error {
	query := `CREATE TABLE IF NOT EXISTS show(
		show_id SERIAL PRIMARY KEY,
		title VARCHAR(255),
		description VARCHAR(1000),
		duration INT,
		show_type VARCHAR(100),
		category_id INT,
		thumbnail_url VARCHAR(255),
		release_date TIMESTAMP,
    	uploader_id INT,
    	CONSTRAINT fk_category_category_id FOREIGN KEY (category_id) REFERENCES category(category_id)
	)`

	_, err := store.db.Exec(query)
	return err
}

func (store *Store) createEpisodeTable() error {
	query := `CREATE TABLE IF NOT EXISTS episode(
		episode_id SERIAL PRIMARY KEY,
		show_id INT,
		number INT,
		title VARCHAR(255),
    	description VARCHAR(1000),
    	video_url VARCHAR(255),
    	thumbnail_url VARCHAR(255),
    	release_date TIMESTAMP,
    	UNIQUE(show_id, number),
    	CONSTRAINT fk_show_show_id FOREIGN KEY (show_id) REFERENCES show(show_id)
	)`

	_, err := store.db.Exec(query)
	return err
}

func (store *Store) createLikeTable() error {
	query := `CREATE TABLE IF NOT EXISTS "like"(
		like_id SERIAL PRIMARY KEY,
		show_id INT,
    	user_id INT,
    	UNIQUE (show_id, user_id),
    	CONSTRAINT fk_show_show_id FOREIGN KEY (show_id) REFERENCES show(show_id)
	)`

	_, err := store.db.Exec(query)
	return err
}

func (store *Store) createFavouriteTable() error {
	query := `CREATE TABLE IF NOT EXISTS favourite(
		favourite_id SERIAL PRIMARY KEY,
		show_id INT,
    	user_id INT,
		UNIQUE(show_id, user_id),
    	CONSTRAINT fk_show_show_id FOREIGN KEY (show_id) REFERENCES show(show_id)
	)`

	_, err := store.db.Exec(query)
	return err
}
