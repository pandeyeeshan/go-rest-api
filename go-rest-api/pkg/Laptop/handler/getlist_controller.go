package handler

import (
	"be-training/go-rest-api/bootstrap"
	"be-training/go-rest-api/pkg/Laptop/model"
	"be-training/go-rest-api/pkg/Laptop/store"
	"net/http"

	"github.com/gin-gonic/gin"
)





func GetList(c *gin.Context) {

	CassandraSession := bootstrap.Session
	var request []model.LaptopStruct


	request = store.Get(CassandraSession)
	
	c.JSON(http.StatusOK, request)

}
