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



// Login godoc
// @Summary      Generate Token
// @Description  get AUTHORIZATION TOKEN
// @Tags         Login
// @Accept       json
// @Produce      json
// @Param       
// @Success      200  {object}  string
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /login [get]
func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {

	group.GET("/login", handler.GenerateToken)
}
