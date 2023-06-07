package handler

import (
	"be-training/go-rest-api/internal/tokenutil"
	"be-training/go-rest-api/pkg/Laptop/model"
	"net/http"

	"github.com/gin-gonic/gin"
)





func GenerateToken(c *gin.Context) {

	// var request []model.LaptopStruct

	
	newUser:=model.User{
		Name :"Eeshan",
		Email : "eeshan@ozone.one",
		Password: "newpassword",
		}
	
		token,err:=tokenutil.CreateAccessToken(&newUser,"hello",913208515142266763)

		
		if(err!=nil){
			panic(err)
		}


		c.JSON(http.StatusOK,gin.H{"token":token})

}
