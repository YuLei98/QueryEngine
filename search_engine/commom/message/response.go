package message

//定义响应数据消息的数据
/*
定义状态：
响应消息：
	nil:server ==> wepserver成功
	其他返回错误string
含数据集：

1.成功返回响应消息
2.其他返回错误

定义类型：

1.状态类型： 返回string 描述
type  ResponseXXX struct{
		stauts:
	}
2.数据类型：
type ResponseXXXData struct {
	xxx
}
type ResponseXXXDatas struct {
	datas []data
	onther
}
*/

type ResponseUpdataStence struct {
	Status string `json:"status"` //响应结果
}

type ResponseStatus struct {
	Stauts string `json:"Stauts"`
}

type RespnseSentenceData struct {
	Url     string  `json:"url"` //图文对
	Caption string  `json:"caption"`
	Score   float64 `json:"score"` //TF-IDF 得分
}
type RespnseSentenceDatas struct {
	Datas []RespnseSentenceData `json:"datas"` //数据量
	Count int                   `json:"count"` //总共的数据量
}

type ResponseBoolMarksData struct {
	Filename string `json:"filename"` //文档名
	Url      string `json:"url"`      //文档url
	Caption  string `json:"caption"`  //文档数据
}

//用户书签管理消息
type ResponseBookMarksDatas struct {
	Count int                     `json:"count"`
	Datas []ResponseBoolMarksData `json:"datas"` //数据量
}

//用户关键词联想和

type ResponseUserkeywordExpansion struct{
	Datas []string   	`json:"Datas"` //联想的数据
}
