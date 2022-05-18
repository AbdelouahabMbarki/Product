package product

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateProduct endpoint.Endpoint
	ListProducts  endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateProduct: makeCreateProductEndpoint(s),
		ListProducts:  makeListProductsEndpoint(s),
	}
}

func makeCreateProductEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateProductRequest)
		ok, err := s.CreateProduct(ctx, req.Name, req.Description, req.Price)
		return CreateProductResponse{Ok: ok}, err
	}
}

func makeListProductsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		products, err := s.ListProducts(ctx)
		return products, err
	}
}
