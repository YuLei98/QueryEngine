package main

import (
	"Search_Engines/search_engines/commom/message"
	"Search_Engines/search_engines/server/engineprocessors"
	"encoding/json"
	"fmt"
)

func main() {
	sev := engineprocessors.New_Search_Engines_Processors()
	datas := message.RequestUpdataStence{
		Username: "ccsu_f",
		Sentence: "中国",
	}
	fp, _ := json.Marshal(datas)
	sev.Request_Message(message.Message{
		Msg:   message.PUT_STR_MES,
		Datas: fp,
	})

	myQry := message.RequestSentenceQry{Username: "ccsu_f", PageNo: 1, Limit: 10}

	fp, _ = json.Marshal(myQry)
	response := sev.Request_Message(message.Message{
		Msg:   message.GET_QRY_STR_MES,
		Datas: fp,
	})
	if response == nil {
		fmt.Println("Error ")
	} else {
		fmt.Println(response)
	}
}
