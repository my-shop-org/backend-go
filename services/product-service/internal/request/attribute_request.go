package request

type AttributeRequest struct {
	Name string `json:"name" validate:"required"`
}

type AttributePatchRequest struct {
	Name *string `json:"name,omitempty"`
}
