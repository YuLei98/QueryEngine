package engineprocessors

import (
	"Search_Engines/search_engines/commom/message"
	"Search_Engines/search_engines/commom/static"
	"encoding/json"
)

/*
功能：用户登录请求
*/
func (this *Search_Engines_Processors) POST_LOGIN_MES(msg message.Message) (resdata interface{}) {

	var request message.RequestUserLogin
	err := json.Unmarshal(msg.Datas, &request)
	if err != nil {
		return message.ResponseStatus{Stauts: "服务器内部错误"}
	}
	resdata = message.ResponseStatus{Stauts: Users.User_Login(static.MY_USER_REDIS, request)}
	return
}

/*
功能：用户注册请求
*/
func (this *Search_Engines_Processors) POST_RIGISTER_MES(msg message.Message) (responseData interface{}) {

	var request message.RequestRegister
	err := json.Unmarshal(msg.Datas, &request)
	if err != nil {
		return err
	}
	err = Users.User_Register(static.MY_USER_REDIS, request)
	if err != nil {
		return message.ResponseStatus{Stauts: err.Error()}
	} else {
		return message.ResponseStatus{Stauts: "success"}
	}
}

/*
功能：用户添加收藏
*/
func (this *Search_Engines_Processors) POST_BOOK_MARKS_MES(msg message.Message) (responseData interface{}) {

	var request message.RequestAddBookMark
	err := json.Unmarshal(msg.Datas, &request)
	if err != nil {
		return message.ResponseStatus{Stauts: err.Error()}
	}

	//fmt.Println("processors 请求POST_BOOK_MARKS_MES ", request)

	err = Users.Add_Book_Mark(request.Username, &request)
	//fmt.Println("processors 完成请求", err)

	if err != nil {
		return message.ResponseStatus{Stauts: err.Error()}
	} else {
		return message.ResponseStatus{Stauts: "success"}
	}

}
