package router

import (
	"voca/internal/ucase"
	"voca/pkg/util"

	"github.com/gin-gonic/gin"
)

func (server *Server) setupRouter() {
	router := gin.Default()
	u := ucase.Ucases{}

	authRoutes := router.Group("/").Use(util.AuthMiddleware(u.TokenMaker))

	router.POST("/users", u.CreateUser)
	router.POST("/users/login", u.LoginUser)

	authRoutes.POST("/accounts", u.CreateAccount)
	authRoutes.GET("/accounts", u.ListAccounts)
	authRoutes.GET("/accounts/:id", u.GetAccountByID)

	authRoutes.POST("/products", u.CreateProduct)
	authRoutes.GET("/products", u.ListProducts)
	authRoutes.GET("/products/:id", u.ListProducts)

	authRoutes.POST("/transfers", u.CreateTransfer)
	server.Router = router

}
