package main

import (
	//"os"

	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/cmd/docs"
	user "github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/User"
	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/people"
	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/projects"
	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/tasks"
	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/teams"

	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/db"
	"github.com/gin-gonic/gin"


	"github.com/gin-contrib/cors"
	"github.com/spf13/viper"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @contact.name   Caio Sousa
// @contact.url    http://www.swagger.io/support
// @contact.email  caiosousafernandesferreira@hotmail.com

// @license.name  BrisaNet 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
    viper.SetConfigFile("./pkg/common/envs/.env")
    viper.ReadInConfig()
    docs.SwaggerInfo.BasePath = "/"

    //port := os.Getenv("PORT") 
    dbUrl := viper.Get("DB_URL").(string)

    r := gin.Default()
    h := db.Init(dbUrl)

    r.Use(cors.Default())
    
    // rotas
    pessoas.RegisterRoutes(r, h)
    projetos.RegisterRoutes(r, h)
    equipes.RegisterRoutes(r, h)
    tasks.RegisterRoutes(r, h)
    user.RegisterRoutes(r, h)
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    //r.Run(":"+port)
    r.Run("localhost:3000")
    //r.Run("localhost:8080")
}