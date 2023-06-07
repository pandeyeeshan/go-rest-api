package handler

import (
	"be-training/go-rest-api/bootstrap"
	"be-training/go-rest-api/pkg/Laptop/model"
	"be-training/go-rest-api/pkg/Laptop/store"
	"net/http"

	"github.com/gin-gonic/gin"
)





func DeleteById(c *gin.Context) {
	CassandraSession := bootstrap.Session

	var request model.IdResponse

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}	



	
	errone:= store.DeleteOne(request.Id,CassandraSession)

	if(errone!=nil){
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
		}

	
	c.JSON(http.StatusOK, model.SuccessResponse{Message: "Deleted"})

}
