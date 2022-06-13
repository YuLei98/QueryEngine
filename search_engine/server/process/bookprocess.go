package process

import (
	"Search_Engines/search_engines/commom/message"
	"Search_Engines/search_engines/server/model"
	"errors"
	"fmt"
	"math"
	"sort"
)

/*
功能：
	enginesprocessor会得到查询到的数据集(文路径) 和 数据的打分
	数据集的分页
	加载指定页面
	返回指定页面
*/
type Book struct {
	Limit     int           //限制大小
	Total     int           //总数据量数据最多的限制
	PageCount int           //总页数
	Count     int           //数据量大小和限制的最小值
	Items     *[]model.Item //数据集
}

/*
功能：获取到数据集排序，分页
参数：无
实现:
1.内部排序算法
2.分页计算算法
返回:无
*/
func (this *Book) Updata_Items() {

	this.Pagination_Cal()
	this.Sort_Item()
	//fmt.Println(this.PageCount)
	//fmt.Println("debug ")
	for i := 0; i < 200 && i < len(*this.Items); i++ {
		//fmt.Println((*this.Items)[i].Score)
	}

}

/*
功能：实现数据集的按打分排序
参数：无
实现:
1.数量少于Limit:快排
2.数量多余Limit:用堆过滤
2.1过滤后更新值并释放空间
返回:无
*/
func (this *Book) Sort_Item() {
	this.Quick_Sort_Items()
}

/*
功能：按分数从高到低排序
参数：无
实现:调用sort.Slice()函数自定义排序
返回：
*/
func (this *Book) Quick_Sort_Items() {
	sort.Slice(*(this.Items), func(i, j int) bool {
		return (*(this.Items))[i].Score > (*(this.Items))[j].Score
	})
}

/*
功能：实现计算总页数
参数：
实现：
1.根据数据集求页面数
2.根据限制求页面数
3.根据1,2取最小值
返回：
*/
func (this *Book) Pagination_Cal() {
	//计算总页数
	// 1.根据数据集求页面数
	pageItems := int(math.Ceil(float64(len(*this.Items)) / float64(this.Limit)))
	//2.根据限制求页面数
	pageCount := int(math.Ceil(float64(this.Total) / float64(this.Limit)))
	//3.根据1,2取最小值
	if pageItems < pageCount {
		this.PageCount = pageItems
	} else {
		this.PageCount = pageCount
	}

	if this.Total < len(*this.Items) {
		this.Count = this.Total
	} else {
		this.Count = len(*this.Items)
	}
	//fmt.Println("func (this *Book) this.Pagination_Cal() {	", this.PageCount)
}

/*
功能：根据页面ID得到文档
参数：无
实现：
1.获取第一个起始页面在文档ID
2.获取页面文档地址
3.Dao层读取文档
4.分配地址---返回响应格式数据
返回：
*/
func (this *Book) Get_Page_By_ID(pageId int) (datas *message.RespnseSentenceDatas, err error) {

	if pageId > this.PageCount {
		return nil, errors.New("page_id 过大")
	}
	//	1.获取第一个起始页面在文档ID
	Docs_Ids := make([]int, 0)
	firstPageId := (pageId - 1) * this.Limit
	itemsLen := len(*this.Items)

	//2.获取页面文档地址
	for i := 0; i < this.Limit; i++ {
		CurPageId := firstPageId + i
		if CurPageId > this.Total || CurPageId >= itemsLen {
			break
		}

		Docs_Ids = append(Docs_Ids, (*(this.Items))[CurPageId].DocId)
	}

	//3.Dao层读取文档
	dao := model.New_Doc_Dao()
	docs, err := dao.Read_File_ID((*this.Items)[firstPageId].RedisId, Docs_Ids...)

	//4.返回响应格式数据
	datas = &message.RespnseSentenceDatas{}

	for i := 0; i < this.Limit; i++ {
		CurPageId := firstPageId + i
		if CurPageId > this.Total || CurPageId >= itemsLen {
			break
		}
		datas.Datas = append(datas.Datas,
			message.RespnseSentenceData{
				Url:     docs[i].Url,
				Caption: docs[i].Document,
				Score:   (*this.Items)[CurPageId].Score})
	}

	datas.Count = this.Count

	fmt.Println( this.Count)
	return
}
func (this *Book) Sample_Test_Show() {
	//fmt.Println(" (this *Book) Sample_Test_Show() ", this.PageCount)

}

/*
功能：返回分页数据
实现：
1.根据limit从新设置分页
2.根据pageId获取数据
3.返回响应数据
*/
func (this *Book) Response_QryResPonseData(sentence string,pageId int, limit int) (datas *message.RespnseSentenceDatas, err error) {
	//1.根据limit从新设置分页
	this.Limit = limit
	this.Pagination_Cal()

	//2.根据pageId获取数据
	datas, err = this.Get_Page_By_ID(pageId)


	
	//3.无数据
	if this.Count == 0 {
		return nil, nil
	}
	//3.返回响应数据
	if err != nil {
		return nil, nil
	}


	//添加2022.6.10
	fwd:=NewWordParticiple()
	keys:=fwd.Participle(sentence)
	render:=&Renderprocess{}
	
	for i:=0;i<len(datas.Datas);i++{
		datas.Datas[i].Caption=render.Render_Data(keys,datas.Datas[i].Caption);
	}
	//end 2022.6.10

	return
}
