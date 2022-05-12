package product

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"content" validate:"required"`
	Price       int32  `json:"updated_at"`
}
