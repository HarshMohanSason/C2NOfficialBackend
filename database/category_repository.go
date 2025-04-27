package database

import (
	"c2nofficialsitebackend/config"
	"c2nofficialsitebackend/models"
	"database/sql"
	"errors"
	"github.com/lib/pq"
)

type CategoryRepository interface {
	AddCategory(product *models.Category) error
	ReturnAllCategories() ([]*models.Category, error)
}
type PostgresCategoryRepository struct {
	DB *sql.DB
}

func (repo *PostgresCategoryRepository) AddCategory(category *models.Category) error {

	query := `INSERT INTO categories (name, size_chart, how_to_measure_image, customization_pdf) VALUES ($1, $2, $3, $4)`
	_, err := repo.DB.Exec(query, category.Name, category.SizeChart, category.HowToMeasureImage, category.CustomizationPDF)

	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return errors.New("category already exists with that name, try another name")
			case "not_null_violation":
				return errors.New(pqErr.Message)
			default:
				return errors.New(pqErr.Message)
			}
		}
		config.LogError(err)
		return errors.New(err.Error())
	}
	return nil
}

func (repo *PostgresCategoryRepository) ReturnAllCategories() ([]*models.CategorySummary, error) {
	query := `SELECT id, name FROM categories`
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*models.CategorySummary
	for rows.Next() {
		var category models.CategorySummary
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}
	return categories, nil
}
