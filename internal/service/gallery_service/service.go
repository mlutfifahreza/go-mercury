package gallery_service

import (
	"go-mercury/internal/data/gallery_db"
)

type Service struct {
	db *gallery_db.DB
}

func NewService(db *gallery_db.DB) Service {
	return Service{db: db}
}

// PRODUCT
func (s Service) GetProductList(filter gallery_db.ProductListFilter) (*gallery_db.ProductList, error) {
	productList, total, err := s.db.GetProducts(filter)
	if err != nil {
		return nil, err
	}
	return &gallery_db.ProductList{
		Products: productList,
		Total:    total,
	}, nil
}

func (s Service) GetProduct(id int) (*gallery_db.Product, error) {
	product, err := s.db.GetProductByID(int64(id))
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s Service) GetProductDetail(id int) (*gallery_db.ProductDetail, error) {
	productDetail, err := s.db.GetProductDetail(int64(id))
	if err != nil {
		return nil, err
	}
	return productDetail, nil
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

// STORE

func (s Service) GetStore(id string) (*gallery_db.Store, error) {
	store, err := s.db.GetStoreByID(id)
	if err != nil {
		return nil, err
	}
	return store, nil
}

func (s Service) DeleteStore(id string) (int, error) {
	affectedCount, err := s.db.DeleteStore(id)
	if err != nil {
		return 0, err
	}
	return affectedCount, nil
}

func (s Service) CreateStore(store gallery_db.Store) error {
	err := s.db.CreateStore(store)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) UpdateStore(store gallery_db.Store) (int, error) {
	affectedCount, err := s.db.UpdateStore(store)
	if err != nil {
		return 0, err
	}
	return affectedCount, nil
}

// LINK

func (s Service) GetLink(id int) (*gallery_db.Link, error) {
	link, err := s.db.GetLinkByID(int64(id))
	if err != nil {
		return nil, err
	}
	return link, nil
}

func (s Service) DeleteLink(id int) (int, error) {
	affectedCount, err := s.db.DeleteLink(int64(id))
	if err != nil {
		return 0, err
	}
	return affectedCount, nil
}

func (s Service) CreateLink(link gallery_db.Link) (int, error) {
	_, err := s.GetProduct(link.ProductId)
	if err != nil {
		return 0, err
	}

	_, err = s.GetStore(link.StoreId)
	if err != nil {
		return 0, err
	}

	id, err := s.db.CreateLink(link)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s Service) UpdateLink(link gallery_db.Link) (int, error) {
	affectedCount, err := s.db.UpdateLink(link)
	if err != nil {
		return 0, err
	}
	return affectedCount, nil
}
