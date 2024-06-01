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
	return product, nil
}

func (s Service) DeleteProduct(id int) (int, error) {
	affectedCount, err := s.db.DeleteProduct(int64(id))
	if err != nil {
		return 0, err
	}
	return affectedCount, nil
}

func (s Service) CreateProduct(product gallery_db.Product) (int, error) {
	id, err := s.db.CreateProduct(product)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s Service) UpdateProduct(product gallery_db.Product) (int, error) {
	affectedCount, err := s.db.UpdateProduct(product)
	if err != nil {
		return 0, err
	}
	return affectedCount, nil
}

func (s Service) GetStore(id int) (*gallery_db.Store, error) {
	store, err := s.db.GetStoreByID(int64(id))
	if err != nil {
		return nil, err
	}
	return store, nil
}

func (s Service) DeleteStore(id int) (int, error) {
	affectedCount, err := s.db.DeleteStore(int64(id))
	if err != nil {
		return 0, err
	}
	return affectedCount, nil
}

func (s Service) CreateStore(store gallery_db.Store) (int, error) {
	id, err := s.db.CreateStore(store)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s Service) UpdateStore(store gallery_db.Store) (int, error) {
	affectedCount, err := s.db.UpdateStore(store)
	if err != nil {
		return 0, err
	}
	return affectedCount, nil
}
