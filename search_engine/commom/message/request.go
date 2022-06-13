package message

//{"username":  "sentence": }
type RequestUpdataStence struct {
	Username string `json:"Username"`
	Sentence string `json:"Sentence"`
}

type RequestSentenceQry struct {
	Username string `json:"Username"`
	PageNo   int    `json:"PageNo"` //页面
	Limit    int    `json:"Limit"` //限制
}

//用户登录
type RequestUserLogin struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type RequestRegister struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
	Passwordc string `json:"Passwordc"`
}

//用户自定义收藏夹
type RequestUPdataBookMark struct {
	Username string `json:"Username"` //文件名
	Filename string `json:"Filename"` //文件名
	Url      string `json:"Url"`      //数据 这里为url-caption
	Caption  string `json:"Caption"`
}

type RequestAddBookMark struct {
	Username string `json:"Username"` //用户名
	Url      string `json:"Url"`      //数据
	Caption  string `json:"Caption"`
}

//用户删除请求
type RequestDelBookMark struct {
	Username string `json:"Username"`
	Filename string `json:"Filename"`
}

//查看用户收藏夹
type RequestQryBookMark struct {
	Username string `json:"Username"`
	PageNo   int    `json:"PageNo"` //页面
	Limit    int    `json:"Limit"` //限制
}


//用户关键词联想
type RequestUserkeywordExpansion struct{
	Username string `json:"Username"`
	InputWord string `json:"InputWord"`
}

//用户敏感词
type RequestEpigraphWords struct{
	Username string `json:"Username"`
	InputWord string `json:"InputWord"`
}