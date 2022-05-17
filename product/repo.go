package product

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-kit/kit/log"
)

var RepoErr = errors.New("Unable to handle Repo Request")

type repo struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (repo *repo) CreateProduct(ctx context.Context, product Product) error {
	sql := `
		INSERT INTO products (id, name, description,price)
		VALUES ($1, $2, $3, $4)`
	if product.Name == "" || product.Description == "" || product.Price == 0 {
		return RepoErr
	}
	_, err := repo.db.ExecContext(ctx, sql, product.ID, product.Name, product.Description, product.Price)
	if err != nil {
		return err
	}
	return nil
}
func (repo *repo) ListProducts(ctx context.Context) ([]Product, error) {
	var products []Product
	rows, err := repo.db.Query("SELECT id, name, description, price FROM products;")
	if err != nil {
		return nil, RepoErr
	}
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price); err != nil {
			return nil, RepoErr
		}
		products = append(products, product)
	}

	return products, nil
}
