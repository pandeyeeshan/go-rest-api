package routes

import (
	"be-training/go-rest-api/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func Returnnew () string {
	return "hello"
}
func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	// ur := repository.NewUserRepository(db, domain.CollectionUser)
	// ur:="Hello"
	
	// group.POST("/login", Returnnew)
}
