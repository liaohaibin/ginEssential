package main

import(
	"github.com/gin-gonic/gin"
	"nti56.com/ginessential/controller"
)

func Route(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/pub/new", controller.CreatePub)
	return r
}