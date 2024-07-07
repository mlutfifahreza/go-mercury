package product_api

const (
	defaultPage = 1
	defaultSize = 10
	maxSize     = 50

	queryPageNumber   = "page_number"
	queryPageSize     = "page_size"
	queryOrderByField = "order_by_field"
	queryOrderByValue = "order_by_value"
)

type OrderByFieldEnum string

const (
	OrderByFieldID OrderByFieldEnum = "id"
)

func (o OrderByFieldEnum) IsValid() bool {
	switch o {
	case OrderByFieldID:
		return true
	default:
		return false
	}
}

type OrderByValueEnum string

const (
	OrderByValueAscending  OrderByValueEnum = "ASC"
	OrderByValueDescending OrderByValueEnum = "DESC"
)

func (o OrderByValueEnum) IsValid() bool {
	switch o {
	case OrderByValueAscending, OrderByValueDescending:
		return true
	default:
		return false
	}
}
