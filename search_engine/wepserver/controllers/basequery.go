package controller

import (
	"Search_Engines/search_engines/commom/message"
	"Search_Engines/search_engines/server/engineprocessors"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

/*
功能：更新用户查询内容
实现：
1.通过参数绑定获取数据
2.server处理数据
3.响应请求
*/
func InterfaceIsNil(i interface{}) bool {
	ret := i == nil
	if !ret {
		defer func() {
			recover()
		}()
		ret = reflect.ValueOf(i).IsNil() // there will exit some panic
	}
	return ret
}
func PUT_STR_MES(c *gin.Context) {
	var request message.RequestUpdataStence
	if err := c.ShouldBind(&request); err == nil {
		fmt.Printf("PUT_STR_MES info:%#v\n", request)

		fp, _ := json.Marshal(request)
		sev := engineprocessors.New_Search_Engines_Processors()
		response := sev.Request_Message(message.Message{
			Msg:   message.PUT_STR_MES,
			Datas: fp,
		})

		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": response,
		})
	}
}
func GET_QRY_STR_MES(c *gin.Context) {

	var request message.RequestSentenceQry
	if err := c.ShouldBind(&request); err == nil {

		//拼request
		username := strings.TrimSpace(c.Param("username"))
		request.Username = username

		fmt.Printf("GET_QRY_STR_MES info:%#v\n", request)
		sev := engineprocessors.New_Search_Engines_Processors()

		//序列化request
		fp, _ := json.Marshal(request)

		response := sev.Request_Message(message.Message{
			Msg:   message.GET_QRY_STR_MES,
			Datas: fp,
		})

		if InterfaceIsNil(response) {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"data": message.GET_QRY_STR_RES_ERROR_MES,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"data": response,
			})
		}
	}
}

func GET_USER_KEY_WORD_EXPANSION(c *gin.Context) {

	var request message.RequestUserkeywordExpansion
	if err := c.ShouldBind(&request); err == nil {

		fmt.Printf("GET_USER_KEY_WORD_EXPANSION info:%#v\n", request)
		sev := engineprocessors.New_Search_Engines_Processors()

		//序列化request
		fp, _ := json.Marshal(request)

		response := sev.Request_Message(message.Message{
			Msg:   message.GET_USER_KEY_WORD_EXPANSION,
			Datas: fp,
		})

		if InterfaceIsNil(response) {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"data": " ",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"data": response,
			})
		}
	}
}
