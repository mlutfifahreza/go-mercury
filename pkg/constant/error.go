package constant

import "errors"

var (
	ErrorInvalidParam = errors.New("invalid_param")

	ProductNotFoundError = errors.New("product_not_found")
	StoreNotFoundError   = errors.New("store_not_found")
)
