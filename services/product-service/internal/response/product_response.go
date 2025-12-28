package response

type ProductResponse struct {
	ID          uint               `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description,omitempty"`
	Categories  []CategoryResponse `json:"categories,omitempty"`
	Price       float64            `json:"price"`
}
