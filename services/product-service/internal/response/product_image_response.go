package response

import "time"

type ProductImageResponse struct {
	ID        uint      `json:"id"`
	ProductID uint      `json:"product_id"`
	VariantID *uint     `json:"variant_id,omitempty"`
	URL       string    `json:"url"`
	IsDefault bool      `json:"is_default"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
