package router

import (
	"fmt"
	"voca/internal/repositories"
	"voca/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	Store      repositories.Store
	TokenMaker util.Maker
	Router     *gin.Engine
	Config     util.Config
}

func NewServer(cfg util.Config, store repositories.Store) (*Server, error) {

	s := util.Config{}
	tokenMaker, err := util.NewPasetoMaker(s.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		Config:     cfg,
		Store:      store,
		TokenMaker: tokenMaker,
	}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", util.GetValidCurrencyValidator())
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) Start(addr string) error {
	return server.Router.Run(addr)
}
