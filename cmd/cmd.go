package cmd

import (
	"main.go/v3/cmd/api"
	"main.go/v3/cmd/database"
	"main.go/v3/cmd/env"
)

func RunAllPackages() {
	env := env.ReadAllEnvs()
	dbs := database.StartDatabases(env)
	api.StartApi(dbs, env)
}
