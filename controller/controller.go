package controller

import (
	"github.com/gin-gonic/gin"
	"myList/models"
	"net/http"
)


func IndexHandler(c *gin.Context){
		c.HTML(http.StatusOK,"index.html",nil)
}

func CreateATodo(c *gin.Context){
	//1.从请求中把数据拿出来
	var todo models.Todo
	c.BindJSON(&todo)
	//2.存入数据库并且返回响应
	err:=models.CreateATodo(&todo)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"code":2000,
			"msg":"success",
			"data" :todo,
		})
	}


}
func GetTodoList(c *gin.Context){
	//查询todo数据表
	todoList,err:=models.GetTodoList()
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"error":err.Error(),
		})
	}else{
		c.JSON(http.StatusOK,todoList)//返回json
	}

}

func UpdateATodo(c *gin.Context){
	id,ok:=c.Params.Get("id")
	if !ok{
		c.JSON(http.StatusOK,gin.H{
			"error":"无效的ID",
		})
		return
	}
	todo,err:=models.GetATodo(id)

	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"error":err.Error(),
		})
		return
	}
	c.BindJSON(&todo)
	if err=models.UpdateATodo(todo);err!=nil{
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
	}else {
		c.JSON(http.StatusOK,todo)
	}

}
func DeleteATodo(c *gin.Context){
	id,ok:=c.Params.Get("id")
	if !ok{
		c.JSON(http.StatusOK,gin.H{"error":"无效的ID"})
		return
	}
	err:=models.DeleteATodo(id)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
	} else {
		c.JSON(http.StatusOK,gin.H{id:"deleted"})
	}
}

