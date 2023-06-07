package routes

import (
	"be-training/go-rest-api/bootstrap"
	"be-training/go-rest-api/pkg/Laptop/handler"
	"time"

	"github.com/gin-gonic/gin"
)

func CRUDoperations(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {


	group.POST("/",handler.CreateList)
	group.GET("/",handler.GetList)
	group.DELETE("/delete",handler.DeleteById)
	group.PUT("/update")

	group.GET("/heartbeat",handler.Heartbeat)


}
