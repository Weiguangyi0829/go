package model

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//在服务器启动后，就初始化一个userCurd实例，把它做成全局的变量
var (
	MyUserCurd *UserCurd
)

type UserCurd struct {
	pool *redis.Pool
}

//使用工厂模式。创建一个UserCurd实例
func NewUserCurd(pool *redis.Pool) (userCurd *UserCurd) {
	userCurd =&UserCurd{
		pool:pool,
	}
	return
}

func (this *UserCurd) getUserCurd(conn redis.Conn,id string) (user *User , err error)  {
	//Query whether the user exists in redis by specifying ID
	res , err :=redis.String(conn.Do("HGet","users","id"))
	if err != nil{
		if err == redis.ErrNil{
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	user = &User{}
	//Unmarshal
	err = json.Unmarshal([]byte(res),user)
	if err != nil{
		fmt.Println("json.Unmarshal err = ",err)
		return
	}
	return
}

//Complete login verification
func (this *UserCurd) Login(userID string , userPwd string) (user *User , err error){
	conn := this.pool.Get()
	defer conn.Close()
	user , err = this.getUserCurd(conn , userID)
	if err != nil{
		return
	}
	//校验密码
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}