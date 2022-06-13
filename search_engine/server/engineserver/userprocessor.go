package engineserver

import (
	"Search_Engines/search_engines/commom/message"
	"Search_Engines/search_engines/server/model"
	"Search_Engines/search_engines/server/process"
	"encoding/json"
	"errors"
)

/*
	面向使用者人群的服务：
*/
type Userprocessor struct {
	bookmarks map[string]*process.Bookmarkprocess //用户收藏夹管理
}

func (this *Userprocessor) Start() {
	this.bookmarks = make(map[string]*process.Bookmarkprocess, 10)
}

/*
功能：改变使用者查询的句子
参数：传入使用者 要改变的句子
实现：
1.查看该使用是否登录
2.改变句子
3.返回改变状态成功：引擎管理者搜索 页面管理者进行跟新存储查询的数据
*/

func (this *Userprocessor) User_Login(redisId int, msg message.RequestUserLogin) (stauts string) {
	fp := &process.Userprocess{}
	stauts, _ = fp.Request_Login(redisId, msg.Username, msg.Password)
	return
}

func (this *Userprocessor) User_Register(redisId int, msg message.RequestRegister) (err error) {

	if msg.Password != msg.Passwordc {
		return errors.New("两次密码不一样")
	}
	fp := &process.Userprocess{}
	err = fp.Request_Register(redisId, msg.Username, msg.Password)
	return
}
func (this *Userprocessor) User_Del(redisId int, username string) (err error) {
	fp := &process.Userprocess{}
	err = fp.Request_Del(redisId, username)
	return
}

/*
功能：添加收藏夹
实现
1.检查是否存在
2.生成写入数据
3.写入数据
*/

func (this *Userprocessor) Try_OK(name string) {
	//1.检查是否存在
	_, ok := this.bookmarks[name]
	if !ok {
		this.bookmarks[name] = &process.Bookmarkprocess{}
		this.bookmarks[name].Init_From_Dir(name)
	}
}

func (this *Userprocessor) Add_Book_Mark(name string, msg *message.RequestAddBookMark) (err error) {

	this.Try_OK(name)

	//2.生成写入数据
	fp := &model.Doc{Url: msg.Url, Document: msg.Caption}
	data, err := json.Marshal(fp)
	if err != nil {
		return err
	}

	//3.写入数据
	err = this.bookmarks[name].Add_Doc(name, data)

	//fmt.Println("userprocessor 完成任务")
	return
}

/*
功能：删除收藏文档
参数：
实现：
1.检查是否存在
2.删除文档
返回值
*/
func (this *Userprocessor) Del_Book_Mark(name string, msg *message.RequestDelBookMark) (err error) {
	this.Try_OK(name)
	err = this.bookmarks[name].Del_Doc(name, msg.Filename)
	return
}

/*
功能：
参数：
实现：
返回值：
*/
func (this *Userprocessor) Qry_Book_Mark(name string, msg *message.RequestQryBookMark) (response *message.ResponseBookMarksDatas) {
	this.Try_OK(name)

	response = this.bookmarks[name].Response(name, msg.PageNo, msg.Limit)
	return
}
func (this *Userprocessor) Updata_Book_Mark(name string, msg *message.RequestUPdataBookMark) (err error) {
	this.Try_OK(name)
	//2.生成写入数据
	fp := &model.Doc{Url: msg.Url, Document: msg.Caption}
	data, err := json.Marshal(fp)
	if err != nil {
		return err
	}
	//更新
	err = this.bookmarks[name].Update_Doc(name, msg.Filename, data)
	return
}
func (this *Userprocessor) End() {

}
