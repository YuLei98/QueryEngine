package engineprocessors

import (
	"Search_Engines/search_engines/commom/message"
	"encoding/json"
	"fmt"
	_"os"
)

func (this *Search_Engines_Processors) PUT_STR_MES(msg message.Message) (resdata interface{}) {
	var request message.RequestUpdataStence
	err := json.Unmarshal(msg.Datas, &request)

	if err != nil {
		return nil
	}

	items, err := Egs.Index_Qry(&request)
	if err != nil {
		return
	}

	err = Books.Update_Qry_Sentence(request.Username, items,request.Sentence)
	if err == nil {
		resdata = message.ResponseUpdataStence{Status: "Success"}
	} else {
		resdata = message.ResponseUpdataStence{Status: "no find"}

	}
	return
}

//用户删除收藏夹数据

func (this *Search_Engines_Processors) PUT_BOOK_MARKS_MES(msg message.Message) (resdata interface{}) {
	var request message.RequestUPdataBookMark
	err := json.Unmarshal(msg.Datas, &request)

	if err != nil {
		return message.ResponseStatus{Stauts: "服务器内部错误"}
	}
	err = Users.Updata_Book_Mark(request.Username, &request)

	if err != nil {
		return message.ResponseStatus{Stauts: "服务器内部错误"}
	} else {
		return message.ResponseStatus{Stauts: "success"}
	}
}

//用户敏感词
func (this *Search_Engines_Processors) PUT_STR_EPIGRAPH_WORD(msg message.Message) (resdata interface{}) {
	var request message.RequestEpigraphWords
	err := json.Unmarshal(msg.Datas, &request)

	if err != nil {
		return message.ResponseStatus{Stauts: "服务器内部错误"}
	}
	
	fmt.Println(request,"成功提交敏感词 ");

	Egs.Updata_epigraph_word(&request);

	if err != nil {
		return message.ResponseStatus{Stauts: "服务器内部错误"}
	} else {
		return message.ResponseStatus{Stauts: "success"}
	}
}
