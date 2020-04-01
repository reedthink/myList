package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin.H 实际上就是 map[string]interface{}
// Response 把后端数据转换成JSON格式并返回前端
func Response(c *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	c.JSON(httpStatus, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

func Success(c *gin.Context, data gin.H, msg string) {
	Response(c, http.StatusOK, 200, data, msg)
}
func Fail(c *gin.Context, data gin.H, msg string) {
	Response(c, http.StatusOK, 400, data, msg)
}
