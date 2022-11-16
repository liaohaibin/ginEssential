package controller

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"nti56.com/ginessential/utils"
	
	"nti56.com/ginessential/common"
	"nti56.com/ginessential/model"
)

func Register(ctx *gin.Context) {
	db := common.GetDB()
	// 获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	// 数据验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号必须为11位",
		})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不少于6位"})
		return
	}
	//如果名称没有传，给一个10位的随机字符串
	if len(name) == 0 {
		name = utils.RangeString(10)
	}
	log.Println(name, telephone, password)
	// 判断手机号是否存在
	if isTelephoneExist(db, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已存在"})
		return
	}

	// 创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	db.Create(&newUser)
	// 返回结果
	ctx.JSON(200, gin.H{
		"message": "注册成功",
	})
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}

	return false
}