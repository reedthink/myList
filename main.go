package main

import (
	"log"
	"os"

	"auth/dao"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	db := dao.InitDB() //初始化数据库
	defer db.Close()   //记得关闭

	r := gin.Default() //新建路由引擎
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
