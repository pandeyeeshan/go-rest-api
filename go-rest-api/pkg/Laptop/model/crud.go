package model

import (
	"context"
)

type LaptopStruct struct {
	Id int `form:"id" json:"id" bson:"_id" `
	Brand string `form:"brand" json:"brand"  bson:"brand"  `
	Model string `form:"model"  json:"model" bson:"model" `
	Year int `form:"year" json:"year" bson:"year" `
	Price int `form:"price" json:"price" bson:"price" `

}

type IdResponse struct {
	Id int `form:"id" json:"id" bson:"id`
}





type CrudUseCase interface {
	Create(c context.Context, user *User) error
	GetUserByEmail(c context.Context, email string) (User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}
