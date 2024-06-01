package gallery_db

type Product struct {
	Id          int64  `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	ImageUrl    string `json:"image_url,omitempty"`
	Description string `json:"description,omitempty"`
}

type Store struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Icon string `json:"icon,omitempty"`
}

type Link struct {
	Id        int    `json:"id"`
	ProductId int    `json:"product_id"`
	StoreId   string `json:"store_id"`
	Link      string `json:"link"`
}
