package product

import "context"

type Product struct {
	ID          string  `json:"id,omitempty"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float32 `json:"price" validate:"required"`
}

type Repository interface {
	CreateProduct(ctx context.Context, product Product) error
	ListProducts(ctx context.Context) ([]Product, error)
}
