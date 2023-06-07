package handler

import (
	"be-training/go-rest-api/bootstrap"
	model "be-training/go-rest-api/pkg/Laptop/model"
	"be-training/go-rest-api/pkg/Laptop/store"
	"net/http"

	"github.com/gin-gonic/gin"
)




func UpdateList(c *gin.Context) {

	CassandraSession := bootstrap.Session
	var request model.Update


	err:=c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}


	errUpdate := store.UpdateOne(request.Id,CassandraSession,request.Brand)
	if errUpdate != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: errUpdate.Error()})
		return
	}
	c.JSON(http.StatusOK,  model.SuccessResponse{Message: "Updated"})

}
