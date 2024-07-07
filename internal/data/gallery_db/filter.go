package gallery_db

type ProductListFilter struct {
	Offset       int
	Limit        int
	OrderByField string
	OrderByValue string
}
