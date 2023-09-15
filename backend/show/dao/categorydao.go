package dao

import (
	"github.com/cyjhuiz/netflix-clone/backend/show/model"
)

type CategoryDao struct {
	store *Store
}

func NewCategoryDao(store *Store) *CategoryDao {
	return &CategoryDao{
		store: store,
	}
}

func (categoryDao *CategoryDao) GetCategories() ([]*model.Category, error) {
	query := `SELECT * FROM category`

	rows, err := categoryDao.store.db.Query(query)
	if err != nil {
		return nil, err
	}

	var categories []*model.Category
	for rows.Next() {
		category := new(model.Category)

		err = rows.Scan(
			&category.CategoryId,
			&category.Name,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (categoryDao *CategoryDao) CreateCategory(category *model.Category) error {
	query := `INSERT INTO category(name) VALUES($1)`

	_, err := categoryDao.store.db.Query(query, category.Name)
	if err != nil {
		return err
	}

	return nil
}
