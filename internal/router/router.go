package router

import (
	"voca/internal/ucase"
	_ "voca/internal/ucase"

	"github.com/gin-gonic/gin"
)

func (server *Server) setupRouter() {
	router := gin.Default()
	u := ucase.Ucases{}

	router.POST("/users", u.CreateAccount)
	router.POST("/users/login", u.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(u.tokenMaker))

	authRoutes.POST("/accounts", u.CreateAccount)
	authRoutes.GET("/accounts", u.listAccounts)
	authRoutes.GET("/accounts/:id", u.getAccountByID)

	authRoutes.POST("/transfers", u.createTransfer)
	server.router = router

}
