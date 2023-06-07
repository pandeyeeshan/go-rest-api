package handler

import (
	"be-training/go-rest-api/bootstrap"
	"be-training/go-rest-api/pkg/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CrudOperationsController struct {
	
}



func CreateList(c *gin.Context) {
	var request model.LaptopStruct

	fmt.Println("1 create list")
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}
	fmt.Println("2 create list")


	err= c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
	}
	fmt.Println("3 create list")

	insert:=bootstrap.CreateDocument(&request)
	fmt.Println(insert)


	 

	

	c.JSON(http.StatusOK, insert)
}


func GetList(c *gin.Context) {

	var request []model.LaptopStruct
	request=bootstrap.Get()
	c.JSON(http.StatusOK, request)

}
