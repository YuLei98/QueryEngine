package engineserver

import (
	"Search_Engines/search_engines/server/model"
	"Search_Engines/search_engines/server/process"
)

/*
页面管理员：
功能：用户更新查询语句后，更新该用户对应的prcoces.Book缓存
*/
type Pageprocessor struct {
	Books map[string]*process.Book
	Sentence map[string]string
}

func (this *Pageprocessor) Start() {
	this.Books = make(map[string]*process.Book, 10)

	this.Sentence=make(map[string]string,10)
}

func (this *Pageprocessor) End() {

}

/*
功能：实现用户查询句子的更新
参数：username:用户名  items:文档集
实现：
1.更换用户查询句子
2.进行分页的预处理
*/
func (this *Pageprocessor) Update_Qry_Sentence(username string, items []model.Item,sentence string) (err error) {
	//1.更换用户查询句子
	//默认最大数据集2e5
	this.Books[username] = &process.Book{Limit: 10,
		Total: 200000,
		Items: &items,
	}
	//2.进行分页的预处理
	this.Books[username].Updata_Items()
	this.Sentence[username]=sentence
	return
}

func (this *Pageprocessor) Response_QryResPonseData(username string, pageNo int, limit int) (datas interface{}, err error) {
	fp, ok := this.Books[username]
	// log.Println(username, "开始查询")
	//fmt.Println(this.Books[username], ok)
	if !ok {
		return nil, nil
	}
	//2022.6.10
	datas, err = fp.Response_QryResPonseData(this.Sentence[username],pageNo, limit)

	return
}
