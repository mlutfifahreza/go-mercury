package product_api

import "go-mercury/internal/data/gallery_db"

type GetProductListRequest struct {
	PageNumber int `json:"page_number,omitempty"`
	PageSize   int `json:"page_size,omitempty"`
}

func (r GetProductListRequest) ConvertToDBFilter() gallery_db.ProductListFilter {
	if r.PageNumber == 0 {
		r.PageNumber = defaultPage
	}

	if r.PageSize == 0 {
		r.PageSize = defaultSize
	} else if r.PageSize > maxSize {
		r.PageSize = maxSize
	}

	return gallery_db.ProductListFilter{
		Offset: (r.PageNumber - 1) * r.PageSize,
		Limit:  r.PageSize,
	}
}

type CreateProductRequest struct {
	Title       string `json:"title" validate:"required,min=3,max=256"`
	ImageUrl    string `json:"image_url" validate:"required,url"`
	Description string `json:"description" validate:"required,min=8,max=512"`
}

type CreateProductResponse struct {
	Id int `json:"id"`
}

type UpdateProductRequest struct {
	Id          int    `json:"id" validate:"required"`
	Title       string `json:"title" validate:"required,min=3,max=256"`
	ImageUrl    string `json:"image_url" validate:"required,url"`
	Description string `json:"description" validate:"required,min=8,max=512"`
}
