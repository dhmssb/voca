package router

import (
	"voca/internal/repositories"
	"voca/pkg/util"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store repositories.Store
	// tokenMaker token.Maker
	router *gin.Engine
	config util.Config
}

func NewServer(cfg util.Config, store repositories.Store) (*Server, error) {

	server := &Server{
		config: cfg,
		store:  store,
		// tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}
