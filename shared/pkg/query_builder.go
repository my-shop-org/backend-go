package pkg

import "gorm.io/gorm"

type QueryBuilder struct {
	query *gorm.DB
}

func NewQueryBuilder(db *gorm.DB) *QueryBuilder {
	return &QueryBuilder{query: db}
}

func (qb *QueryBuilder) Build() *gorm.DB {
	return qb.query
}

func (qb *QueryBuilder) ApplyPagination(limit, offset int) *QueryBuilder {
	qb.query = qb.query.Limit(limit).Offset(offset)
	return qb
}

func (qb *QueryBuilder) ApplySorting(sortBy string, ascending bool) *QueryBuilder {
	if sortBy != "" {
		var order string
		if ascending {
			order = "ASC"
		} else {
			order = "DESC"
		}

		qb.query = qb.query.Order(sortBy + " " + order)
	}

	return qb
}

func (qb *QueryBuilder) ApplyFilter(filter map[string]interface{}) *QueryBuilder {
	for key, value := range filter {
		qb.query = qb.query.Where(key+" = ?", value)
	}
	return qb
}

func (qb *QueryBuilder) ApplyPreload(relations []string) *QueryBuilder {
	for _, relation := range relations {
		qb.query = qb.query.Preload(relation)
	}
	return qb
}
