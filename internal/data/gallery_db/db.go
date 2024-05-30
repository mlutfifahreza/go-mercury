package gallery_db

type DB struct {
	host     string
	port     int
	username string
	password string
}

func (db DB) GetProduct(id int) (*Product, error) {
	return &Product{
		ID:       id,
		ImageUrl: "some-url.com/image.jpg",
		Title:    "some title",
		Currency: "IDR",
		Price:    123000,
	}, nil
}
