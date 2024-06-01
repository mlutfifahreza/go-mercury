package gallery_db

type Product struct {
	Id          int64  `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	ImageUrl    string `json:"image_url,omitempty"`
	Description string `json:"description,omitempty"`
}

type Store struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Icon string `json:"icon,omitempty"`
}
