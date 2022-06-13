package engineprocessors

import (
	"Search_Engines/search_engines/commom/message"
	"Search_Engines/search_engines/server/engineserver"
	"Search_Engines/search_engines/server/model"
	"Search_Engines/search_engines/server/process"
	"fmt"
	"os"
	"time"
)

var (
	Egs   *engineserver.EngineProcessor
	Users *engineserver.Userprocessor
	Books *engineserver.Pageprocessor
)

/*
 功能:开发过程中简单测试docDao功能s
 参数：无
 实现：
 //flushall清空数据库
 1.初始哈redis链接池
 2.服务启动每个docDao实例
*/
var (
	dao       *model.DocDao
	enginedao *model.IndexEngineDao
)

type Search_Engines_Processors struct {
	engineprocessor *engineserver.EngineProcessor
	userprocessor   *engineserver.Userprocessor
	books           *engineserver.Pageprocessor
}

/*
功能：返回一个引擎服务管理值
*/
func New_Search_Engines_Processors() (egs *Search_Engines_Processors) {
	return &Search_Engines_Processors{
		engineprocessor: Egs,
		userprocessor:   Users,
		books:           Books,
	}
}

//优雅开机
func init() {
	Start()

	Egs = &engineserver.EngineProcessor{}
	Users = &engineserver.Userprocessor{}
	Books = &engineserver.Pageprocessor{}

	Users.Start()
	Books.Start()
	start :=time.Now().Round(time.Millisecond).UnixNano() / 1e9

	err := Egs.Start(1, 0)
	if err != nil {
		fmt.Println("开机失败！！！")
		os.Exit(0)
	} else {
		fmt.Println("开机成功！！！go go go")
		fmt.Println("包含索引数量 ",len(Egs.IndexProcess.Tree.Keys()))
		fmt.Println(" 文档数量 ",Egs.IndexProcess.Keys);
	}
	end:=time.Now().Round(time.Millisecond).UnixNano() / 1e9
	fmt.Printf("开机耗时秒：%v", end - start)
	
}
func Start() {
	//当服务器启动是初始化redis的链接池
	InitPool("localhost:6379", 20, 0, 300*time.Second)
	dao = model.New_Init_DocDao(POOL)
	enginedao = model.New_Init_Index_Engine_Dao(POOL)
	model.New_Init_UserDao(POOL)
	process.NewWordParticiple()
}

//优雅关机
func End() {
	// Egs.End()
	Users.End()
	Books.End()
}

/*
 功能：实现处理消息数据接收返回处理后的数据
 实现：
 	1.根据commom/message/中消息进行断言
	2.消息处理 然后给数据给不同管理者处理
	2.拼装返回消息
*/
func (this *Search_Engines_Processors) Request_Message(msg message.Message) (responsedata interface{}) {

	if msg.Msg == message.PUT_STR_MES {
		//查询句子更新
		responsedata = this.PUT_STR_MES(msg)

	} else if msg.Msg == message.GET_QRY_STR_MES {
		//获取查询句子
		responsedata = this.GET_QRY_STR_MES(msg)

	} else if msg.Msg == message.POST_LOGIN_MES {
		//请求登录
		responsedata = this.POST_LOGIN_MES(msg)

	} else if msg.Msg == message.POST_RIGISTER_MES {
		//请求注册
		responsedata = this.POST_RIGISTER_MES(msg)

	} else if msg.Msg == message.GET_BOOK_MARKS_MES {
		//请求用户收藏夹
		responsedata = this.GET_BOOK_MARKS_MES(msg)

	} else if msg.Msg == message.POST_BOOK_MARKS_MES {
		//请求用户添加收藏
		responsedata = this.POST_BOOK_MARKS_MES(msg)

	} else if msg.Msg == message.PUT_BOOK_MARKS_MES {
		//编辑用户收藏夹消息
		responsedata = this.PUT_BOOK_MARKS_MES(msg)

	} else if msg.Msg == message.DELETE_BOOK_MARKS_MES {
		//用户删除收藏夹数据
		responsedata = this.DELETE_BOOK_MARKS_MES(msg)
	} else if msg.Msg == message.GET_USER_KEY_WORD_EXPANSION {
		//用户关键词联想
		responsedata = this.GET_USER_KEY_WORD_EXPANSION(msg)

	}else if msg.Msg==message.PUT_STR_EPIGRAPH_WORD{
		//用户敏感词
		responsedata=this.PUT_STR_EPIGRAPH_WORD(msg);
	}

	//fmt.Println("sever服务返回 ", responsedata)

	return
}
