package request

type ProductRequest struct {
	Name         string  `json:"name" validate:"required"`
	Description  string  `json:"description,omitempty"`
	Categories   []uint  `json:"categories" validate:"required,min=1"`
	BasePrice    float64 `json:"base_price" validate:"required,gt=0"`
	ComparePrice float64 `json:"compare_price" `
}

type ProductPatchRequest struct {
	Name         *string  `json:"name,omitempty"`
	Description  *string  `json:"description,omitempty"`
	Categories   *[]uint  `json:"categories,omitempty"`
	BasePrice    *float64 `json:"base_price,omitempty" validate:"omitempty,gt=0"`
	ComparePrice *float64 `json:"compare_price,omitempty" validate:"omitempty,gt=0"`
}
