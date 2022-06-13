package engineprocessors

import (
	"Search_Engines/search_engines/commom/message"
	"encoding/json"
)

/*
功能：用户删除收藏夹文档
参数：
实现：
返回
*/
func (this *Search_Engines_Processors) DELETE_BOOK_MARKS_MES(msg message.Message) (responseData interface{}) {

	var request message.RequestDelBookMark
	err := json.Unmarshal(msg.Datas, &request)

	if err != nil {
		return nil
	}

	err = Users.Del_Book_Mark(request.Username, &request)
	if err == nil {
		return message.ResponseStatus{Stauts: "success"}
	} else {
		return message.ResponseStatus{Stauts: err.Error()}
	}
}
