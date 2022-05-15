package product

import (
	"context"
	"encoding/json"
	"net/http"
)

type (
	CreateProductRequest struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float32 `json:"price"`
	}
	CreateProductResponse struct {
		Ok string `json:"ok"`
	}

	ListProductsRequest struct {
	}
	ListProductsResponse struct {
		Products []Product `json:"products"`
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
func decodeUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateProductRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
