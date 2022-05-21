package product

import (
	"context"
	"encoding/json"
	"net/http"
)

type (
	CreateProductRequest struct {
		Name        string  `json:"name"`
		Sku         string  `json:"sku"`
		Description string  `json:"description"`
		Price       float32 `json:"price"`
	}
	CreateProductResponse struct {
		Ok string `json:"ok"`
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

func decodeUserListProductsReq(ctx context.Context, r *http.Request) (request interface{}, err error) {
	return r, nil
}
