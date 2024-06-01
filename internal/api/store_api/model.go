package store_api

type CreateStoreRequest struct {
	Name string `json:"name" validate:"required,min=3,max=256"`
	Icon string `json:"icon" validate:"required,url"`
}

type CreateStoreResponse struct {
	Id int `json:"id"`
}

type UpdateStoreRequest struct {
	Id   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required,min=3,max=256"`
	Icon string `json:"icon" validate:"required,url"`
}
