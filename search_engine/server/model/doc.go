package model

import "fmt"

type Doc struct {
	//确定字段 要保证反序列成功
	Url      string `json:"url"`     //文档链接
	Document string `json:"caption"` //文档内容
}

/*
功能：debug的展示
参数：无
实现：打印Url+Document以，号分割
返回：无
*/
func (this *Doc) Debug_Show() {
	fmt.Println(this.Url, ",", this.Document)
}
