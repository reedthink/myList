package main

import (
	"fmt"
	ginSwagger "github.com/swaggo/gin-swagger"
	"myList/controller"
	_ "myList/docs"
	"myList/middleware"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	r.Use(m1) //全局注册中间件函数
	r.GET("/list", controller.IndexHandler)
	r.Use(middleware.AuthMiddleware())
	//API v1
	v1Group := r.Group("v1")
	{ //少见的换行花括号
		//todoList
		//添加
		v1Group.POST("/todo", controller.CreateATodo)
		//查看所有待办事项
		v1Group.GET("/todo", controller.GetTodoList)
		//修改
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		//删除
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}

//中间件,统计耗时
func m1(c *gin.Context) {
	fmt.Println("m1 in...")
	//计时
	start := time.Now()
	c.Next() //调用后续的处理函数
	//c.Abort() //阻止调用
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Printf("m1 out.\n")
}
