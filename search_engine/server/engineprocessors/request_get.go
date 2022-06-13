package engineprocessors

import (
	"Search_Engines/search_engines/commom/message"
	"encoding/json"
)

/*
功能：
参数：
实现：
1.pagepr
返回
*/
func (this *Search_Engines_Processors) GET_QRY_STR_MES(msg message.Message) (responseData interface{}) {

	var request message.RequestSentenceQry
	err := json.Unmarshal(msg.Datas, &request)

	if err != nil {
		return nil
	}
	data, err := Books.Response_QryResPonseData(request.Username, request.PageNo, request.Limit)
	if err != nil {
		return nil
	}
	responseData = data
	return
}

//查看用户收藏夹
func (this *Search_Engines_Processors) GET_BOOK_MARKS_MES(msg message.Message) (responseData interface{}) {

	var request message.RequestQryBookMark
	err := json.Unmarshal(msg.Datas, &request)

	if err != nil {
		return nil
	}
	//fmt.Println("GET_BOOK_MARKS_MES 请求", request)
	responseData = Users.Qry_Book_Mark(request.Username, &request)
	//fmt.Println("GET_BOOK_MARKS_MES 返回", responseData)

	return
}

//用户关键词联想
func (this *Search_Engines_Processors) GET_USER_KEY_WORD_EXPANSION(msg message.Message) (responseData interface{}) {

	var request message.RequestUserkeywordExpansion
	err := json.Unmarshal(msg.Datas, &request)

	if err != nil {
		return nil
	}
	responseData = Egs.Request_Word_KEY_WORD_EXPANSION(&request)
	return
}
