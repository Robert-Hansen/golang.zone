package repositories

import (
	"log"

	"database/sql"

	"errors"
	"github.com/steffen25/golang.zone/database"
	"github.com/steffen25/golang.zone/models"
)

type CategoryRepository interface {
	Create(c *models.Category) error
	GetAll() ([]*models.Category, error)
	FindById(id int) (*models.Category, error)
	FindByName(name string) (*models.Category, error)
	Exists(name string) bool
	Delete(id int) error
	Update(c *models.Category) error
}

type categoryRepository struct {
	*database.MySQLDB
}

func NewCategoryRepository(db *database.MySQLDB) CategoryRepository {
	return &categoryRepository{db}
}

// Create a category.
func (cr *categoryRepository) Create(c *models.Category) error {
	exists := cr.Exists(c.Name)
	if exists {
		return errors.New("category name already exists")
	}

	stmt, err := cr.DB.Prepare("INSERT INTO categories SET name=?, created_at=?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(c.Name, c.CreatedAt.Format("20060102150405"))
	if err != nil {
		return err
	}

	cId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	c.ID = int(cId)

	return nil
}

// GetAll retrieves all categories.
func (cr *categoryRepository) GetAll() ([]*models.Category, error) {
	var categories []*models.Category

	rows, err := cr.DB.Query(`SELECT id,
 										   name,
 										   created_at,
 										   updated_at
 										   FROM categories`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		c := new(models.Category)
		err := rows.Scan(&c.ID,
			&c.Name,
			&c.CreatedAt,
			&c.UpdatedAt)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

// FindById retrieves a category by its ID.
func (cr *categoryRepository) FindById(id int) (*models.Category, error) {
	category := models.Category{}

	err := cr.DB.QueryRow(`SELECT c.id,
 										c.name,
 										c.created_at,
 										c.updated_at
 										FROM categories c
 										WHERE c.id = ?`,
		id).Scan(&category.ID,
		&category.Name,
		&category.CreatedAt,
		&category.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

// FindByName retrieves a category by its Name.
func (cr *categoryRepository) FindByName(name string) (*models.Category, error) {
	category := models.Category{}
	err := cr.DB.QueryRow(`SELECT c.id,
 										c.name,
										c.created_at,
										c.updated_at
										FROM categories c
										WHERE c.name
										LIKE ?`,
		"%"+name+"%").Scan(
		&category.ID,
		&category.Name,
		&category.CreatedAt,
		&category.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &category, nil
}

// Delete a category by its ID.
func (cr *categoryRepository) Delete(id int) error {
	return nil
}

// Update a category.
func (cr *categoryRepository) Update(c *models.Category) error {
	exists := cr.Exists(c.Name)
	if exists {
		return errors.New("Category does already exists")
	}

	var categoryId int
	err := cr.DB.QueryRow(`SELECT id
								 FROM categories
								 WHERE id = ?`, c.ID).Scan(&categoryId)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if c.ID == categoryId {
		err := cr.updateCategory(c)
		if err != nil {
			return err
		}

		return nil
	}

	err = cr.updateCategory(c)
	if err != nil {
		return err
	}

	return nil
}

// Exists check if category already exists.
func (cr *categoryRepository) Exists(name string) bool {
	var exists bool
	err := cr.DB.QueryRow(`SELECT EXISTS
								(SELECT id
								 FROM categories
								 WHERE name=?)`, name).Scan(&exists)
	if err != nil {
		log.Printf("[CATEGORY REPO]: Exists err %v", err)
		return true
	}

	return exists
}

// updateCategory
func (cr *categoryRepository) updateCategory(c *models.Category) error {
	stmt, err := cr.DB.Prepare(`UPDATE categories
 									  SET name = ?,
 									   	  updated_at = ?
 									  WHERE id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(c.Name, c.UpdatedAt, c.ID)
	if err != nil {
		return err
	}

	return nil
}
