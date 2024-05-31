package gallery_service

import "go-mercury/internal/data/gallery_db"

type Service struct {
	db *gallery_db.DB
}

func NewService(db *gallery_db.DB) Service {
	return Service{db: db}
}

func (s Service) GetProduct(id int) (*gallery_db.Product, error) {
	product, err := s.db.GetProductByID(int64(id))
	if err != nil {
		return nil, err
	}
	return &product, nil
}