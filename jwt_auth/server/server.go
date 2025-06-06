package server

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type BasicServer interface {
	Run()
	RegisterBasicRoutes()
}

type basicServer struct {
	engine *gin.Engine
	db     *mongo.Database
}

func NewBasicServer(r *gin.Engine, db *mongo.Database) BasicServer {
	return &basicServer{r, db}
}

func (s *basicServer) Run() {
	s.engine.Run("localhost:8881")
}
