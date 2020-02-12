package routes

import (
	"github.com/gin-gonic/gin"
	"main.go/v3/cmd/api/routes/register"
	"main.go/v3/cmd/database"
)

type Service struct {
	router *gin.Engine
	dbs    *database.AllDatabases
}

func RunAllRoutes(r *gin.Engine, dbs *database.AllDatabases) *Service {
	s := &Service{}
	s.router = r
	s.dbs = dbs
	s.registerGroup()
	return s
}

func (s *Service) registerGroup() {
	repo := register.Service(s.router, s.dbs)
	registerGroup := s.router.Group("/register")

	registerGroup.GET("/add", repo.AddRegister())

}
