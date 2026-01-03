package params

type BaseQueryParam struct {
	Limit  int `query:"limit"`
	Offset int `query:"offset"`
}

type ProductQueryParam struct {
	BaseQueryParam
	CategoryID string `query:"category_id"`
}

func NewProductQueryParam() *ProductQueryParam {
	return &ProductQueryParam{
		BaseQueryParam: BaseQueryParam{
			Limit:  10,
			Offset: 0,
		},
	}
}

type VariantQueryParam struct {
	BaseQueryParam
	ProductID string `query:"product_id"`
}

func NewVariantQueryParam() *VariantQueryParam {
	return &VariantQueryParam{
		BaseQueryParam: BaseQueryParam{
			Limit:  10,
			Offset: 0,
		},
	}
}
