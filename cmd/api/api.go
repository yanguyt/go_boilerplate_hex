package api

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"main.go/v3/cmd/api/routes"
	"main.go/v3/cmd/database"
	"main.go/v3/cmd/env"
)

type ApiRoutes interface {
}

type AppAPI struct {
	Router *gin.Engine
	Dbs    *database.AllDatabases
}

func StartApi(dbs *database.AllDatabases, env *env.AllEnvs) *AppAPI {
	app := &AppAPI{}

	if env.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.Use(gin.Recovery())
	app.Router = router
	app.Dbs = dbs

	routes.RunAllRoutes(router, dbs)

	color.Green("Api Rodando na porta :%v", env.Port)
	app.Router.Run(fmt.Sprintf(":%v", env.Port))

	return app
}
