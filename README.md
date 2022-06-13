# 搜索引擎

# 快速开始

## 开发环境

```
开发环境：( 腾讯云 免费 服务器 一个 ) + ( vscode ) + ( Google Chrome 网络浏览器 )
xxx@VM-12-4-ubuntu:~/src/Search_Engines/search_engines$ uname -a
Linux VM-12-4-ubuntu 5.4.0-109-generic #123-Ubuntu SMP Fri Apr 8 09:10:54 UTC 2022 x86_64 x86_64 x86_64 GNU/Linux
```
测试数据：悟空数据集 中的 一百万 条 数据
```go
加载自定义词典!!!1
加载自定义词典成功 ACEPT 001!!!!
开机成功！！！go go go
包含索引数量  1341927
 文档数量  1175173
开机耗时秒：42 cpu核心数: 2
[GIN-debug] [WARNING] Now Gin requires Go 1.14+.
```
## [项目演示](http://121.5.218.2:8080/)

## 基础功能：
```
纯文本 信息存储 & 搜索
搜索 结果 分页 展示
自定义 关键词 过滤
相关 搜索 功能
```
## 拓展功能
```
支持 用户 注册，登录
用户 可以收藏 搜索结果 到 个人 收藏夹
用户 可以删除 个人收藏夹 内容
```

# api文档

## 基础功能
### 远程关机(#关机后必须得在服务器上开机)
```
接口地址	/shoutdown
请求方式	Any
表单参数	无
```
### 搜索提示
```
接口地址   /search/tips
请求方式   GET
表单参数 Username,InputWord
```
### 提交敏感词
```
接口地址  /wu_kong_search/Epigraph_words
请求方式  PUT
表单参数 Username,InputWord
```
###  搜索句子
```
1.更新查询句子

接口地址 /wu_kong_search/updata
接口方式 PUT
表单参数 Username,Sentence

2.返回查询分页

接口地址 /user/search/:username
接口方式 PUT
表单参数 Username,PageNo,Limit
```

## 用户管理
### 用户登录
```
接口地址  /user
请求方式  POST
表单参数 Username,Password
```
### 用户注册
```
接口地址  /register
请求方式  POST
表单参数 Username,Password,Passwordc
```
### 收藏夹 数据 获取
```
接口地址  /dir/get/:username
请求方式  POST
表单参数 Username,Password,Passwordc
```
### 收藏夹 数据 添加
```
接口地址 /dir/add
请求方式 POST
表单参数 Username,Url,Caption
```
### 收藏夹 数据 删除
```
接口地址 /dir/del
请求方式 POST
表单参数 Username,Filename
```
### 收藏夹 数据 修改
```
接口地址 /dir/
请求方式 PUT
表单参数  Username,Url,Caption
```
# 技术栈总结文档

## 项目结构
```
├── commom    wep服务和后端服务接口
│   ├── message 存放定义的消息 一个请求一个响应
│   ├── static  全局的静态文件
│   └── utils   wep服务和后端公用的工具类
├── server  后端服务者
│   ├── engineprocessors 后端服务的管理者
│   ├── engineserver    后端各个服务管理人
│   ├── model           后端与数据库对接层
│   ├── process         后端处理器
│   ├── storagesystem   添加的存储系统
│   ├── test            开发过程中简单测试的地方
│   └── untils          后台工具
└── wepserver   前端服务者
    ├── controllers 服务器
    │   └── basequery.go
    ├── logs    日志文件
    ├── main    添加路由和api定义
    │   ├── API.md
    │   └── main.go 
    ├── static 存放前端所有静态文件目录
    └── views   html所在文件目录
```
## 索引设计
### 索引建立
![索引建立](https://git.acwing.com/ccsu_f/search_engines/-/blob/master/imgs/index00.png)
![索引描述](https://git.acwing.com/ccsu_f/search_engines/-/blob/master/imgs/index01.png)
![索引作业](https://git.acwing.com/ccsu_f/search_engines/-/blob/master/imgs/index02.png)
### 参考实现
```go
package process
import (
	"Search_Engines/search_engines/server/model"
	"Search_Engines/search_engines/server/untils"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
	avl "github.com/emirpasic/gods/trees/avltree"
)

type Iddexprocess struct {
	sync.RWMutex
	Tree    *avl.Tree `json:"tree"`    //平衡二叉树 数据库ID
	RedisID int       `json:"redisId"` //数据库的ID
	Keys    uint64    `json:"keys"`    //当前数据库的生成key值 开始为0
}

/*
功能： 给一个索引插入文档ID
实现： 1.avl树中有该索引,直接该索引对应val中的值插入
	   2.avl树中中不存在该索引,索引中插入一个集合，并将该doc_id放入
参数: key：关键字, doc_id文档ID
返回值 返回是否成功
*/
func (this *Iddexprocess) Insert(key interface{}, doc_Id int) { //插入一个索引

	tf, ok := this.Tree.Get(key)
	if ok {
		fp, ok := tf.(map[int]int) //这个map的key值存文档ID 第二个value存该索引在该文档的得分
		if !ok {
			return
		}
		fp[doc_Id] = fp[doc_Id] + 1
	} else {
		fp := make(map[int]int, 1)
		fp[doc_Id] = 1
		this.Tree.Put(key, fp)
	}
}
/*
功能:根据索引得到文档ID
实现：
1.根据key找到对应value
2.生成untils.set找索引
参数: key
返回值:实现的set 列表 err信息
*/
func (this *Iddexprocess) Get_Doc_Ids(key interface{}) (uids *untils.Set, err error) {

	node := this.Tree.GetNode(key)
	if node != nil {
		tf := node.Value
		fp, ok := tf.(map[int]int)
		if ok {
			//实现 map[int]int 转 map[int]bool
			fpc := make(map[int]bool)
			for k := range fp {
				fpc[k] = true
			}
			uids = untils.New_From_Map(fpc)
			return
		}
	} else { //没有索引
		err = errors.New("no")
		return
	}
	err = errors.New("fail ")
	return
}
/*
功能:增加一篇文档
实现：
1.生成唯一ID
2.分词(url+ 句子分词)
3.构建索引
4.文档缓存数据库
参数:doc:文档
返回值:
*/
func (this *Iddexprocess) Add_Doc(doc *model.Doc) (err error) {
	//考虑线程安全
	this.Lock()
	defer this.Unlock()

	//1.生成唯一ID
	doc_id := int(this.Keys)
	this.Keys = this.Keys + 1 //关键字加一
	log.Println("写", this.Keys)
	//2.句子分词不过滤
	strs := Wdp.Participle_No_Filtering(doc.Document)

	//3.构建索引
	for _, v := range strs {
		this.Insert(v, doc_id)
	}

	dDao := model.New_Doc_Dao()

	//4.2文档缓存磁盘
	err = dDao.Write_File_Doc(this.RedisID, doc_id, doc)
	if err != nil {
		log.Println("文档缓存错误!!!")
		this.Keys = this.Keys - 1
		return
	}

	return
}

/*
功能:查询根据索引句子查询
参数：str string: 查询句 ,info string：敏感词
实现:
1.去重分词
2.根据索引查文档ID
3.合并文档ID
4.返回文档ID集合
5.实现数据处理打分
	5.1.遍历词
	5.2.遍历文档
	5.3.获取打分
返回：文档ID集合
*/
func (this *Iddexprocess) Index_Qry(str string,info string) (result []model.Item, err error) {


	//1.分词
	strs := Wdp.Participle_Filter(str,info);
	//2.查文档ID
	uids := untils.New() //
	flag := true         //定义开始

	fmt.Println("开始查询", strs)
	start := time.Now().UnixNano()

	//3.合并文档ID


	mytools := &untils.MyTools{}
	for i := 0; i < len(strs); i++ {

		v, err := this.Get_Doc_Ids(strs[i])

		if mytools.InterfaceIsNil(v) == true {

			return nil, errors.New("No find!!1")
		}
		if err != nil {
			return nil, errors.New("No find")
		}
		if flag == true {
			uids = v
			flag = false
		} else {
			uids = uids.Intersect(v)
		}

	}

	if flag == true {
		return nil, errors.New("No find")
	}
	fmt.Printf("查询耗时毫秒：%v", (time.Now().UnixNano()-start)/1000000)
	//4.返回文档ID集合
	fmt.Println("开始打分", str)
	start = time.Now().UnixNano()

	scorerprcoss := &ScorerPrcoss{}
	var numdocs uint64 = this.Keys
	docs_id := uids.List() //，tf:表示词条在文档d中出现的频率，这个比较简单。
	docsIdLen := len(docs_id)
	//计算每篇文档得分(每个关键词典的TF-IDF分值和)
	// TF 表示词条在文档d中出现的频率，这个比较简单。
	// IDF可以由文档库中的总文档数（numDocs）除以包含该词条的文档数量（docFreq），再将得到的商取以10为底的对数得到，即：

	//5.实现数据处理打分

	//5.1.遍历词
	//5.2.遍历文档
	//5.3.获取打分
	ans := make(map[int]float64, docsIdLen)

	for _, v := range strs {

		node, ok := this.Tree.Get(v)

		if !ok {
			continue
		}

		fp, ok := node.(map[int]int) //这个map的key值存文档ID 第二个value存该索引在该文档的得分

		docfreq := len(fp) //包含该词条的文档数量（docFreq）
		for _, doc_id := range docs_id {
			TF, ok := fp[doc_id] //val:这个文档出现v这个分词次数
			if !ok {
				continue
			}
			//这个文档中有这个关键词
			ans[doc_id] = ans[doc_id] + scorerprcoss.TF_IDF(uint32(TF), numdocs, uint64(docfreq))
		}
	}

	fmt.Printf("打分耗时毫秒: %v\n", (time.Now().UnixNano()-start)/1000000)
	//返回数据
	for k, v := range ans {
		result = append(result, model.Item{RedisId: this.RedisID, DocId: k, Score: v})
	}
	return
}

```
## 打分算法
### TF-IDF介绍
![tf-idf介绍](https://git.acwing.com/ccsu_f/search_engines/-/blob/master/imgs/doc00.png)
### TF-IDF实现
```go
package process

import "math"

/*
 用于打分算法处理器
*/
type ScorerPrcoss struct{}

/*
功能：实现TF-IDF打分
参数：
实现:
TF: 词频(Term Frequency)，表示词条在文档d中出现的频率，这个比较简单。
IDF可以由文档库中的总文档数（numDocs）除以包含该词条的文档数量（docFreq），再将得到的商取以10为底的对数得到，即：
返回：分数
*/
func (this *ScorerPrcoss) TF_IDF(tf uint32, numDocs uint64, docFreq uint64) (Score float64) {
	idf := math.Log10(float64(numDocs) / float64(docFreq+1))
	return float64(idf) * float64(tf)
}
```
## 数据分页

### 分页介绍
![分页介绍](https://git.acwing.com/ccsu_f/search_engines/-/blob/master/imgs/pa00.png)
### 参考实现
```go
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


	return
}
```
## 性能优化
### go的并发
![如何跑满服务器](https://git.acwing.com/ccsu_f/search_engines/-/blob/master/imgs/go00.png)
![性能优化思路](https://git.acwing.com/ccsu_f/search_engines/-/blob/master/imgs/go01.png)
### 参考实现
```go
package main

import (
	"fmt"
	"runtime"
	"time"
)
type Ad_test struct{ XX int }

type qry struct{}

func (this *qry) get(a int) (result int) {
	time.Sleep(time.Second * 1)
	return a * 2
}

type diaodu struct{}

//向myChan中写入数据
func (slef *diaodu) putNum(sum int, adChan chan Ad_test) {

	for i := 1; i <= sum; i++ {
		adChan <- Ad_test{XX: i}
	}
	//关闭 myChan
	close(adChan)
}

func (slef *diaodu) get(adChan chan Ad_test, intChan chan int, exitChan chan bool) {
	//使用for循环
	//var
	for {
		num, ok := <-adChan
		if !ok { //如果取不到值
			break
		}
		fp := &qry{}
		res := fp.get(num.XX)
		intChan <- res
	}
	//fmt.Println("有一个协程取不数据退出")
	//这里不关闭primeChan
	exitChan <- true

}

func (self *diaodu) huoqu(sum int, peoples int) (result []int) {
	//解决协程中出现panic
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error ")
		}
	}()

	adChan := make(chan Ad_test, sum)
	intChan := make(chan int, sum) //记录答案的管道
	exitChan := make(chan bool, peoples)
	ad := &diaodu{}
	go ad.putNum(sum, adChan)
	for i := 0; i < peoples; i++ {
		go ad.get(adChan, intChan, exitChan)
	}
	for i := 0; i < peoples; i++ {
		<-exitChan
	}
	close(intChan)
	for v := range intChan {
		result = append(result, v)
	}
	return
}
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	
	start := time.Now().UnixNano()
	fp := diaodu{}
	fp.huoqu(10, 10)
	fmt.Println("结束 所有  协程")
	fmt.Println("耗时 毫秒 \n", (time.Now().UnixNano()-start)/1000000)

}

```
# 资料和引用
```
gofound-Golang实现的 全文 搜索引擎: https://github.com/newpanjing/gofound
jieba分词：https://github.com/fxsjy/jieba
搜索引擎 基本 技术  原理: https://developer.aliyun.com/article/765914
跳表：https://blog.csdn.net/yjw123456/article/details/105159817/
```
![谢谢 结束！](https://git.acwing.com/ccsu_f/search_engines/-/blob/master/imgs/goend.png)
