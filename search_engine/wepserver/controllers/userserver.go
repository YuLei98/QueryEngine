package controller

import (
	"Search_Engines/search_engines/commom/message"
	"Search_Engines/search_engines/server/engineprocessors"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
功能:用户登录
实现：
1.通过参数绑定获取数据
2.server处理数据
3.响应消息
*/

func POST_LOGIN_MES(c *gin.Context) {
	var request message.RequestUserLogin
	if err := c.ShouldBind(&request); err == nil {
		//拼request
		fmt.Printf("POST_LOGIN_MES info:%#v\n", request)

		sev := engineprocessors.New_Search_Engines_Processors()
		//序列化request
		fp, _ := json.Marshal(request)

		response := sev.Request_Message(message.Message{
			Msg:   message.POST_LOGIN_MES,
			Datas: fp,
		})

		fmt.Println("登录返回", response)

		if InterfaceIsNil(response) {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"data": "sucess",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"data": response,
			})
		}

	}
}

func POST_RIGISTER_MES(c *gin.Context) {
	var request message.RequestRegister
	if err := c.ShouldBind(&request); err == nil {
		//拼request
		fmt.Printf("POST_RIGISTER_MES info:%#v\n", request)
		sev := engineprocessors.New_Search_Engines_Processors()
		//序列化request
		fp, _ := json.Marshal(request)

		response := sev.Request_Message(message.Message{
			Msg:   message.POST_RIGISTER_MES,
			Datas: fp,
		})

		if InterfaceIsNil(response) {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"data": "sucess",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"data": response,
			})
		}

	}
}


func PUT_STR_EPIGRAPH_WORD(c *gin.Context) {
	var request message.RequestEpigraphWords
	if err := c.ShouldBind(&request); err == nil {
		//拼request
		fmt.Printf("PUT_STR_EPIGRAPH_WORD info:%#v\n", request)
		sev := engineprocessors.New_Search_Engines_Processors()
		//序列化request
		fp, _ := json.Marshal(request)

		response := sev.Request_Message(message.Message{
			Msg:   message.PUT_STR_EPIGRAPH_WORD,
			Datas: fp,
		})

		if InterfaceIsNil(response) {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"data": "sucess",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"data": response,
			})
		}

	}
}


