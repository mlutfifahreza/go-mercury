package gallery_service

type LinkDetail struct {
	ID        int    `json:"id,omitempty"`
	ProductID int    `json:"product_id,omitempty"`
	Link      string `json:"link,omitempty"`
	StoreName string `json:"store_name,omitempty"`
	StoreIcon string `json:"store_icon,omitempty"`
}

type ProductDetail struct {
	ID          int          `json:"id,omitempty"`
	Title       string       `json:"title,omitempty"`
	ImageURL    string       `json:"image_url,omitempty"`
	Description string       `json:"description,omitempty"`
	LinkDetails []LinkDetail `json:"link_details,omitempty"`
}
