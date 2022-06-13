package model

//定义一个用户结构体

type User struct {
	//确定字段 要保证反序列成功
	UserName string `json:"userName"` //用户名
	UserPwd  string `json:"userPwd"`  //用户密码
}
