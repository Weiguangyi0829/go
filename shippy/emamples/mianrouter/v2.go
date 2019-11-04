package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	V "shippy/emamples/mianrouter/view"

)

func Auth() gin.HandlerFunc  {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if len(token) == 0{

		}
	}
}

func x() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Request.FormValue("username")
		re := regexp.MustCompile("[~!;@#$%^&*(){}|<>\\\\/+\\-【】:\"?'：；‘’“”，。、《》\\]\\[`]")
		isIllegal := re.MatchString(username)
		if isIllegal  {
			c.JSON(http.StatusBadRequest,gin.H{
				"status": "请重新输入",
			})
			c.Abort()// 如果有特殊字符 就终止请求向下走
			return
		}

	}
}
//[~!@#$%^&*(){}|<>\\\\/+\\-【】:\"?'：；‘’“”，。、《》\\]\\[`]
func main() {
	router := gin.Default()

	router.POST("/register",x(), V.S2)
	router.POST("/login",x(), V.S3)

	router.POST("/test", x(),func(c *gin.Context) {
		username := c.PostForm("username")
		fmt.Println("you input right username", username)
		c.JSON(http.StatusOK, gin.H{
			"status": "S",
		})
	})

	router.Run(":9000")
}
