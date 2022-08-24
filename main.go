package main

import (
	"os"

	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/docs"
	user "github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/User"
	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/people"
	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/projects"
	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/tasks"
	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/teams"

	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/db"
	"github.com/gin-gonic/gin"



	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @contact.name   Caio Sousa
// @contact.url    http://www.swagger.io/support
// @contact.email  caiosousafernandesferreira@hotmail.com

// @license.name  Mozilla Public License 2.0
// @license.url   https://www.mozilla.org/en-US/MPL/2.0/
// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func main() {
	

	docs.SwaggerInfo.Title = "Gerenciador de Projetos"
	docs.SwaggerInfo.Description = "REST API Desafio"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "golang-posgre-brisanet.herokuapp.com"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"https"}

	port := os.Getenv("PORT")
	dbUrl := os.Getenv("DATABASE_URL")

	r := gin.Default()
	h := db.Init(dbUrl)

	
	r.Use(CORSMiddleware())

	// rotas
	pessoas.RegisterRoutes(r, h)
	projetos.RegisterRoutes(r, h)
	equipes.RegisterRoutes(r, h)
	tasks.RegisterRoutes(r, h)
	user.RegisterRoutes(r, h)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":" + port)
	//r.Run("localhost:3000")
	//export PATH=$(go env GOPATH)/bin:$PATH
	
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Headers", "*")
        
            c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
            c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
            c.Writer.Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
            c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, DELETE, OPTIONS, PATCH")
        

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}