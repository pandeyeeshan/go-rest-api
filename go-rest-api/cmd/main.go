package main

import (
	"be-training/go-rest-api/bootstrap"
	"be-training/go-rest-api/pkg/routes"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	app := bootstrap.App()

	env := app.Env
	// bootstrap.NewCassandraDatabase(env)
	CassandraSession := bootstrap.Session
	fmt.Println(CassandraSession)
  	// defer CassandraSession.Close()

	// defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	routes.Setup(env, timeout, gin)

	gin.Run(env.ServerAddress)
}
