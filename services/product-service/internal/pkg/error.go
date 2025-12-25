package pkg

import "errors"

var (
	ParentCategoryNotFound       = errors.New("parent category not found")
	CategoryNotFound             = errors.New("category not found")
	CategoryCannotBeItsOwnParent = errors.New("category cannot be its own parent")
	CategoryHasChildren          = errors.New("category has child categories and cannot be deleted")
)
