package main

import (
	"github.com/gin-gonic/gin"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/go-sql-driver/mysql"
	"nti56.com/ginessential/common"
	"github.com/spf13/viper"
	"os"
)


func main() {
	InitConfig()
	db := common.InitDB()
	defer db.Close()
	r := gin.Default()
	r = Route(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	r.Run()
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {

	}
}

