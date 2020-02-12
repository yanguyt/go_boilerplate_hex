package database

import (
	"main.go/v3/cmd/env"
)

//AllDatabases is responsible to manage all instances of Databases
type AllDatabases struct {
}

//StartDatabases is responsible to fill all database infos
func StartDatabases(e *env.AllEnvs) *AllDatabases {
	dbs := &AllDatabases{}

	return dbs
}
