package main

import (
	"be-training/go-rest-api/bootstrap"
	"be-training/go-rest-api/cmd/docs"
	"be-training/go-rest-api/pkg/Laptop/routes"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	app := bootstrap.App()

	env := app.Env
	// bootstrap.NewCassandraDatabase(env)
	CassandraSession := bootstrap.Session
	fmt.Println(CassandraSession)
  	defer CassandraSession.Close()

	// defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	routes.Setup(env, timeout, gin)

	gin.Run(env.ServerAddress)
}
