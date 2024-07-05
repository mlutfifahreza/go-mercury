package store_api

type CreateStoreRequest struct {
	Id    string `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required,min=3,max=256"`
	Icon  string `json:"icon" validate:"required,url"`
	Color string `json:"color" validate:"required"`
}

type CreateStoreResponse struct {
	Id string `json:"id"`
}

type UpdateStoreRequest struct {
	Id    string `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required,min=3,max=256"`
	Icon  string `json:"icon" validate:"required,url"`
	Color string `json:"color" validate:"required"`
}
