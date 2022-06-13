package controller

import (
	"Search_Engines/search_engines/commom/message"
	"Search_Engines/search_engines/server/engineprocessors"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//查询用户收藏夹
func GET_BOOK_MARKS_MES(c *gin.Context) {

	var request message.RequestQryBookMark
	if err := c.ShouldBind(&request); err == nil {
		//拼request
		fmt.Printf("RequestQryBookMark info:%#v\n", request)
		//拼request
		username := strings.TrimSpace(c.Param("username"))
		request.Username = username
		

		sev := engineprocessors.New_Search_Engines_Processors()
		//序列化request
		fp, _ := json.Marshal(request)

		response := sev.Request_Message(message.Message{
			Msg:   message.GET_BOOK_MARKS_MES,
			Datas: fp,
		})

		//fmt.Println("登录返回", response)

		if InterfaceIsNil(response) {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"data": "NO FIND",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"data": response,
			})
		}

	}
}

//用户添加收藏
func POST_BOOK_MARKS_MES(c *gin.Context) {

	var request message.RequestAddBookMark
	if err := c.ShouldBind(&request); err == nil {
		//拼request
		fmt.Printf("POST_BOOK_MARKS_MES info:%#v\n", request)

		sev := engineprocessors.New_Search_Engines_Processors()
		//序列化request
		fp, _ := json.Marshal(request)

		response := sev.Request_Message(message.Message{
			Msg:   message.POST_BOOK_MARKS_MES,
			Datas: fp,
		})

		fmt.Println("登录返回", response)

		if InterfaceIsNil(response) {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"data": "NO FIND",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"data": response,
			})
		}

	}
}

//编辑用户收藏夹消息
func PUT_BOOK_MARKS_MES(c *gin.Context) {

	var request message.RequestUPdataBookMark
	if err := c.ShouldBind(&request); err == nil {
		//拼request
		fmt.Printf("PUT_BOOK_MARKS_MES info:%#v\n", request)

		sev := engineprocessors.New_Search_Engines_Processors()
		//序列化request
		fp, _ := json.Marshal(request)

		response := sev.Request_Message(message.Message{
			Msg:   message.PUT_BOOK_MARKS_MES,
			Datas: fp,
		})

		fmt.Println("登录返回", response)

		if InterfaceIsNil(response) {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"data": "NO FIND",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"data": response,
			})
		}

	}
}

//用户删除收藏夹数据
func DELETE_BOOK_MARKS_MES(c *gin.Context) {

	var request message.RequestDelBookMark
	if err := c.ShouldBind(&request); err == nil {
		//拼request
		fmt.Printf("DELETE_BOOK_MARKS_MES info:%#v\n", request)

		sev := engineprocessors.New_Search_Engines_Processors()
		//序列化request
		fp, _ := json.Marshal(request)

		response := sev.Request_Message(message.Message{
			Msg:   message.DELETE_BOOK_MARKS_MES,
			Datas: fp,
		})

		fmt.Println("登录返回", response)

		if InterfaceIsNil(response) {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"data": "NO FIND",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"data": response,
			})
		}

	}
}
