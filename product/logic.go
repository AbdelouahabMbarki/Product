package product

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
)

type service struct {
	repository Repository
	logger     log.Logger
}

func NewService(rep Repository, logger log.Logger) Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func (s service) CreateProduct(ctx context.Context, name string, description string, price float32) (string, error) {
	logger := log.With(s.logger, "method", "CreateProduct")
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	product := Product{
		ID:          id,
		Name:        name,
		Description: description,
		Price:       price,
	}
	if err := s.repository.CreateProduct(ctx, product); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}
	logger.Log("create product", id)

	return "Success", nil
}

func (s service) ListProducts(ctx context.Context) ([]Product, error) {
	all := "All"
	logger := log.With(s.logger, "method", "ListProducts")

	products, err := s.repository.ListProducts(ctx)

	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}

	logger.Log("Listing Products", all)

	return products, nil
}
