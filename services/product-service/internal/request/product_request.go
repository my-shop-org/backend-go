package request

type ProductRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description,omitempty"`
	Categories  []uint  `json:"categories" validate:"required,min=1"`
	Price       float64 `json:"price" validate:"required,gt=0"`
}

type ProductPatchRequest struct {
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Categories  *[]uint  `json:"categories,omitempty"`
	Price       *float64 `json:"price,omitempty" validate:"omitempty,gt=0"`
}
