package link_api

type CreateLinkRequest struct {
	ProductId int    `json:"product_id" validate:"required"`
	StoreId   string `json:"store_id" validate:"required"`
	Link      string `json:"link" validate:"required,url"`
}

type CreateLinkResponse struct {
	Id int `json:"id"`
}

type UpdateLinkRequest struct {
	Id        int    `json:"id" validate:"required"`
	ProductId int    `json:"product_id" validate:"required"`
	StoreId   string `json:"store_id" validate:"required"`
	Link      string `json:"link" validate:"required,url"`
}
