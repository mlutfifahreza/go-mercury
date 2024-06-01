package gallery_db

type Product struct {
	ID          int64  `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	ImageURL    string `json:"image_url,omitempty"`
	Description string `json:"description,omitempty"`
}
