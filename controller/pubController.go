package controller

import (
	// "log"
	// "net/http"
	"github.com/gin-gonic/gin"
	// "github.com/jinzhu/gorm"

	// "nti56.com/ginessential/utils"
	"nti56.com/ginessential/common"
	"nti56.com/ginessential/model"

)

func CreatePub(ctx *gin.Context) {
	db := common.GetDB()
	user := ctx.PostForm("user")
	project := ctx.PostForm("project")
	destHost := ctx.PostForm("desthost")

	newPub := model.Pub {
		User: user,
		Project: project,
		DestHost: destHost,
	}
	db.Create(&newPub)
	ctx.JSON(200, gin.H{
		"message": "新发布了",
	})
}