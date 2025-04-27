package database

import (
	"c2nofficialsitebackend/config"
	"c2nofficialsitebackend/models"
	"database/sql"
	"errors"
	"github.com/lib/pq"
)

type ProductRepository interface {
	AddProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
	DeleteProduct(product *models.Product) error
	SearchProduct(product *models.Product) (*models.Product, error)
	DeleteAllProducts() error
}

type PostgresProductRepository struct {
	DB *sql.DB
}

func (repo *PostgresProductRepository) AddProduct(product *models.Product) error {
	query := `
		INSERT INTO products (
			name, category_id, long_description, short_description, thumbnail_image,
			carousel_images, slug, price, discount, inventory, sku, status,
			weight, width, length, height
		)
		VALUES (
			$1, $2, $3, $4, $5,
			$6, $7, $8, $9, $10, $11, $12,
			$13, $14, $15, $16
		)`
	_, err := repo.DB.Exec(query,
		product.Name,
		product.Category.ID,
		product.LongDescription,
		product.ShortDescription,
		product.ThumbnailImage,
		pq.Array(product.CarouselImages), // Wrapping arrays
		product.Slug,
		product.Price,
		product.Discount,
		product.Inventory,
		product.SKU,
		product.Status,
		product.Weight,
		product.Width,
		product.Length,
		product.Height,
	)

	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return errors.New(pqErr.Message)
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

func (repo *PostgresProductRepository) UpdateProduct(product *models.Product) error {

	return nil
}

func (repo *PostgresProductRepository) DeleteProduct(product *models.Product) error {

	return nil
}

func (repo *PostgresProductRepository) DeleteAllProducts() error {

	return nil
}

func (repo *PostgresProductRepository) SearchProduct(product *models.Product) (*models.Product, error) {

	return nil, nil
}
