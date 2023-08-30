package server

import (
	"backend/db"
	"backend/server/middlewares"
	"backend/server/panier"
	"github.com/gin-gonic/gin"
)

type Server struct {
	*gin.Engine
}

func New(db *db.Client) *Server {
	r := gin.Default()

	r.Use(middlewares.GenerateRequestID())
	r.GET("/healthcheck", healthcheck)

	// Add db instance
	panier.AddEndpoints(r.RouterGroup, db)

	return &Server{
		Engine: r,
	}
}
