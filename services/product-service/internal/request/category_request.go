package request

import "github.com/google/uuid"

type CategoryRequest struct {
	Name        string     `json:"name" validate:"required"`
	Description string     `json:"description,omitempty"`
	ParentID    *uuid.UUID `json:"parent_id,omitempty"`
}
