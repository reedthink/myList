package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"myList/dao"
)
// @title myList API
// @version 0.1
// @description 一个todo清单后端，支持多用户注册，登录，每个用户有独立的清单
func main() {
	InitConfig()
	db := dao.InitDB() //初始化数据库
	defer db.Close()   //记得关闭

	r := gin.Default()   //新建路由引擎
	r = CollectRouter(r) //路由组

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
		log.Printf("没找到配置文件")
	}
}
