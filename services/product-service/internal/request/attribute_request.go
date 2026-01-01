package request

type AttributeRequest struct {
	Name string `json:"name" validate:"required"`
}

type AttributePatchRequest struct {
	Name *string `json:"name,omitempty"`
}

type AttributeValueRequest struct {
	AttributeID uint   `json:"attribute_id" validate:"required"`
	Value       string `json:"value" validate:"required"`
}

type AttributeValuePatchRequest struct {
	AttributeID *uint   `json:"attribute_id,omitempty"`
	Value       *string `json:"value,omitempty"`
}
