package model

type ErrorResponse struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}


type Update struct {
	Id int `json:"id"`
	Brand string `json:"brand"`
}