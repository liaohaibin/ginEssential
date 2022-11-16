package main

import (
	"github.com/gin-gonic/gin"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/go-sql-driver/mysql"
	"nti56.com/ginessential/common"
)


func main() {
	db := common.InitDB()
	defer db.Close()
	r := gin.Default()
	r = Route(r)
	panic(r.Run())

}



