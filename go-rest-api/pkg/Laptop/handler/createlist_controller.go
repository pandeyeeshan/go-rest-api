package handler

import (
	"be-training/go-rest-api/bootstrap"
	"be-training/go-rest-api/pkg/Laptop/model"
	"be-training/go-rest-api/pkg/Laptop/store"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


func CreateList(c *gin.Context) {
	CassandraSession := bootstrap.Session

	var request model.LaptopStruct

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}


	insert:=store.CreateDocument(CassandraSession,&request)

	fmt.Print(insert)	 

	

	c.JSON(http.StatusCreated, model.SuccessResponse{Message: "Created"})
}
