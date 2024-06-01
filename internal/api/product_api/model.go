package product_api

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
