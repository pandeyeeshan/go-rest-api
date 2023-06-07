package routes

import (
	"be-training/go-rest-api/bootstrap"
	"be-training/go-rest-api/pkg/Laptop/handler"
	"time"

	"github.com/gin-gonic/gin"
)

func Returnnew () string {
	return "hello"
}
func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {

	group.GET("/login", handler.GenerateToken)
}
