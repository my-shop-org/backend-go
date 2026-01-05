package request

type VariantRequest struct {
	ProductID       uint                 `json:"product_id" validate:"required"`
	SKU             string               `json:"sku" validate:"required"`
	BasePrice       float64              `json:"base_price" validate:"required,gt=0"`
	ComparePrice    float64              `json:"compare_price"`
	Stock           int                  `json:"stock" validate:"gte=0"`
	AttributeValues []uint               `json:"attribute_values"`
	ProductImages   []*ProductImageInput `json:"product_images,omitempty"`
}

type VariantPatchRequest struct {
	ProductID       *uint                `json:"product_id,omitempty"`
	SKU             *string              `json:"sku,omitempty"`
	BasePrice       *float64             `json:"base_price,omitempty" validate:"omitempty,gt=0"`
	ComparePrice    *float64             `json:"compare_price,omitempty" validate:"omitempty,gte=0"`
	Stock           *int                 `json:"stock,omitempty" validate:"omitempty,gte=0"`
	AttributeValues *[]uint              `json:"attribute_values,omitempty"`
	ProductImages   *[]*ProductImageInput `json:"product_images,omitempty"`
}
