package message

//确定消息类型
const  (
	LoginMesType  = "LoginMes"
	LoginResMesType = "LoginResMes"
)

type Message struct {
	Type string `json:"type"`
	Date string	 `json:"date"`
}

//定义2个消息
type LoginMes struct {
	UserID string `json:"user_id"`
	UserPwd string `json:"user_pwd"`
	Username string `json:"username"`
}

type LoginResMes struct {
	Code int  `json:"code"`// 500表示未注册  200 表示登录成功
	Error string `json:"error"`
}