package routes

import (
	"time"

	"be-training/go-rest-api/bootstrap"
	"be-training/go-rest-api/pkg/Laptop/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, gin *gin.Engine) {

  
	publicRouter := gin.Group("")
	// All Public APIs

	NewLoginRouter(env, timeout, publicRouter)

	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	CRUDoperations(env, timeout, protectedRouter)


}
