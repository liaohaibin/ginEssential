package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"nti56.com/ginessential/model"
)

var DB *gorm.DB
func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "10.1.20.198"
	port := "32306"
	database := "ginessential"
	username := "root"
	password := "Nti56@com"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset,
	)

	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Pub{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}