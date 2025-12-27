package response

type CategoryResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	ParentID    *uint  `json:"parent_id,omitempty"`
}

type CategoryTreeResponse struct {
	CategoryResponse
	Children []*CategoryTreeResponse `json:"children,omitempty"`
}
