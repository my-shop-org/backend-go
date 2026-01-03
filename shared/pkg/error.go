package pkg

import "errors"

var (
	ParentCategoryNotFound         = errors.New("parent category not found")
	CategoryNotFound               = errors.New("category not found")
	CategoryCannotBeItsOwnParent   = errors.New("category cannot be its own parent")
	CategoryHasChildren            = errors.New("category has child categories and cannot be deleted")
	ProductNotFound                = errors.New("product not found")
	DuplicateEntry                 = errors.New("Duplicated entry found.")
	NoFieldsToUpdate               = errors.New("no fields provided to update")
	AttributeNotFound              = errors.New("attribute not found")
	AttributeHasValues             = errors.New("attribute has values and cannot be deleted")
	AttributeValueNotFound         = errors.New("attribute value not found")
	VariantNotFound                = errors.New("variant not found")
	InvalidAttributeValueForProduct = errors.New("attribute value does not belong to product's attributes")
)
