package routes

import (
	"be-training/go-rest-api/bootstrap"
	"be-training/go-rest-api/pkg/handler"
	"time"

	"github.com/gin-gonic/gin"
)

func CRUDoperationsEndoint(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {


	group.POST("/create",handler.CreateList)
	group.GET("/heartbeat",handler.Heartbeat)
	group.GET("/fetch",handler.GetList)
	group.PUT("/update")
	group.DELETE("/delete")

}
