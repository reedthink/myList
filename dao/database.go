package dao

import (
	"fmt"

	"github.com/spf13/viper"

	"auth/model"

	"github.com/jinzhu/gorm"
)
//dao层负责数据库连接并且绑定
var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := viper.GetString("datasource.driverName") //viper是自动读取制定文件的吗？
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True",
		username,
		password,
		host,
		port,
		database,
		charset,
	)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("连接数据库失败！" + err.Error())
	}
	db.AutoMigrate(&model.User{})//自动绑定
	DB = db
	return db
}
func GetDB() *gorm.DB {
	return DB
}
