package response

type AttributeValueResponse struct {
	ID    uint   `json:"id"`
	Value string `json:"value"`
}

type AttributeResponse struct {
	ID     uint                     `json:"id"`
	Name   string                   `json:"name"`
	Values []AttributeValueResponse `json:"values,omitempty"`
}

type AttributeValueDetailResponse struct {
	ID          uint   `json:"id"`
	AttributeID uint   `json:"attribute_id"`
	Value       string `json:"value"`
}
