package gallery_service

type Product struct {
	ID       int    `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	ImageUrl string `json:"imageUrl,omitempty"`
	Currency string `json:"currency,omitempty"`
	Price    int    `json:"price,omitempty"`
}
