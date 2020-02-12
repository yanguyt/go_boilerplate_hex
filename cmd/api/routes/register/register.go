package register

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/v3/cmd/database"
)

type Repo struct {
	dbs    *database.AllDatabases
	router *gin.Engine
}

func Service(r *gin.Engine, dbs *database.AllDatabases) *Repo {
	repo := &Repo{}

	repo.dbs = dbs
	repo.router = r

	return repo
}

func (r *Repo) AddRegister() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"teste": "deu bom"})
	}
}
