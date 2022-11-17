package main

import(
	"github.com/gin-gonic/gin"
	"nti56.com/ginessential/controller"
	"nti56.com/ginessential/middlewares"

)

func Route(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/pub/new", controller.CreatePub)
	r.POST("/api/auth/login", controller.Login)
	r.POST("/api/auth/info", middlewares.AuthMiddleware(), controller.Info)
	return r
}