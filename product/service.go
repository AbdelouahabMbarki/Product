package product

import "context"

type Service interface {
	CreateProduct(ctx context.Context, name string, sku string, description string, price float32) (string, error)
	ListProducts(ctx context.Context) ([]Product, error)
}
