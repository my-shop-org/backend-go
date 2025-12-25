package response

type CategoryResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	ParentID    *string `json:"parent_id,omitempty"`
}

type CategoryTreeResponse struct {
	CategoryResponse
	Children []*CategoryTreeResponse `json:"children,omitempty"`
}
