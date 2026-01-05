package request

type ProductImageRequest struct {
	ProductID  uint   `json:"product_id" validate:"required"`
	VariantID  *uint  `json:"variant_id,omitempty"`
	URL        string `json:"url" validate:"required,url"`
	IsDefault  bool   `json:"is_default"`
}

type ProductImagePatchRequest struct {
	VariantID  *uint   `json:"variant_id,omitempty"`
	URL        *string `json:"url,omitempty" validate:"omitempty,url"`
	IsDefault  *bool   `json:"is_default,omitempty"`
}
