package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func InitMySQL() {
	driverName := viper.GetString("datasource.driverName") //viper是自动读取制定文件的吗？
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	//dsn:="root:sqldb@tcp(127.0.0.1:3306)/myList?charset=utf8mb4&parseTime=True&loc=Local"
	//todo 时区会有问题，暂时搁置 ，因为此数据库中未储存时间
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset,
	)

	db, err := gorm.Open(driverName, dsn) //使用全局变量，不用冒号

	if err != nil {
		panic(err) //连不上数据库，属于异常情况，而不是错误
	}
	DB = db
	log.Println("数据库成功连接")
}
