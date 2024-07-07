package gallery_db

type Product struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	ImageUrl    string `json:"image_url"`
	Description string `json:"description"`
}

type ProductList struct {
	Products []Product `json:"products,omitempty"`
	Total    int       `json:"total,omitempty"`
}

type Store struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Icon  string `json:"icon"`
	Color string `json:"color"`
}

type Link struct {
	Id        int    `json:"id"`
	ProductId int    `json:"product_id"`
	StoreId   string `json:"store_id"`
	Link      string `json:"link"`
}

type LinkDetail struct {
	ID         int    `json:"id"`
	Link       string `json:"link"`
	StoreName  string `json:"store_name"`
	StoreIcon  string `json:"store_icon"`
	StoreColor string `json:"store_color"`
}

type ProductDetail struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	ImageURL    string       `json:"image_url"`
	Description string       `json:"description"`
	LinkDetails []LinkDetail `json:"link_details"`
}
