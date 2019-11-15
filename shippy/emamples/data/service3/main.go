package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
	"log"
	"net/http"
	"regexp"
	"shippy/emamples/service3/athu/model/authentication"
	V "shippy/emamples/service3/view"
	"time"
)

type S struct {}

func Auth() gin.HandlerFunc  {
	return func(c *gin.Context) {
		username := c.Request.FormValue("username")
		password := c.Request.FormValue("password")
		z := &authentication.User{
			Username: username,
			Password: password,
		}
		d := time.Duration(1000)*time.Second
		token := c.Request.Header.Get("token")
		fmt.Println(token)
		if len(token) == 0{
			c.JSON(http.StatusUnauthorized,gin.H{
				"status":"you hava no token",
			})
			c.Abort()
			return
		}
		_, err := authentication.ParseToken(token)
		if err != nil{
			token, _ = authentication.JwtGenerateToken(z,d)
			c.JSON(http.StatusUnauthorized,gin.H{
				"status":"your token is wrong",
				"token":token,
			})
			c.Abort()
		}
		c.Next()
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
			c.Abort()
		}

	}
}

func (s *S) Anything(c *gin.Context) {
	log.Print("Received Say.Anything API request")
	c.JSON(200, map[string]string{
		"message": "Hi, this is the Greeter API",
	})
}

func main()  {
	service := web.NewService(
			web.Name("go.micro.api.user"),
	)
	service.Init()
	//O.NewUserService("go.micro.Usrv.greeter",client.DefaultClient)
	s := new(S)
	router := gin.Default()

	router.POST("/user/login",Auth(),x(),V.S3)
	router.POST("/user/register",x(),V.S2)
	router.POST("/user",x(), s.Anything)

	service.Handle("/",router)
	if err := service.Run(); err != nil{
		log.Fatal(err)
	}

}