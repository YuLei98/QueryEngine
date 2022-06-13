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
		// fmt.Printf("%T \n", tf)
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

	//fmt.Println(doc)

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

	// fmt.Println(this.Tree)
	// os.Exit(0)

	dDao := model.New_Doc_Dao()
	//4.1文档缓存redis数据库
	// // err = dDao.Write_Doc(this.RedisID, doc_id, doc)
	// // if err != nil {
	// // 	fmt.Println("文档缓存错误!!!")
	// // 	this.Keys = this.Keys - 1
	// // 	return
	// // }

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
参数：传入要查询的句子
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
	//fmt.Println(strs)

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





func (this *Iddexprocess) Delete_Doc() {

}

/*
功能：进行打分多线程优化
实现：对应每篇文档遍历词通过管道
*/
func (this *Iddexprocess) Multi_Process_Index_TF_IDF_adChan(list []int, adChan chan int) {
	//解决协程中出现的panic
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error ")
		}
	}()

	for _, v := range list {
		adChan <- v
	}
	close(adChan)
}
func (this *Iddexprocess) Multi_Process_Index_TF_IDF_resultChan(strs []string, ans *map[int]float64, adChan chan int, exitChan chan bool) {
	scorerprcoss := &ScorerPrcoss{}
	var numdocs uint64 = this.Keys

	for {
		doc_id, ok := <-adChan
		if !ok { //取不到值
			break
		}
		for _, v := range strs {
			node, ok := this.Tree.Get(v)
			if !ok {
				continue
			}
			fp, ok := node.(map[int]int) //这个map的key值存文档ID 第二个value存该索引在该文档的得分
			if !ok {
				continue
			}
			docfreq := len(fp)   //包含该词条的文档数量（docFreq）
			TF, ok := fp[doc_id] //val:这个文档出现v这个分词次数
			if !ok {
				continue
			}

			//这个文档中有这个关键词
			this.Lock()
			(*ans)[doc_id] = (*ans)[doc_id] + scorerprcoss.TF_IDF(uint32(TF), numdocs, uint64(docfreq))
			this.Unlock()
		}

	}
	exitChan <- true

}
func (this *Iddexprocess) Multi_Process_Index_TF_IDF(strs []string, uids *untils.Set) (result []model.Item, err error) {
	//特判：
	if uids.Count() == 0 {
		return nil, errors.New("空")
	}
	//1.文档集合
	docs := uids.List()
	docsIdLen := len(docs)
	peoples := 10

	ans := make(map[int]float64, docsIdLen)

	//2.文档集
	adChan := make(chan int, docsIdLen)
	exitChan := make(chan bool, peoples)

	go this.Multi_Process_Index_TF_IDF_adChan(docs, adChan)

	for i := 0; i < peoples; i++ {
		go this.Multi_Process_Index_TF_IDF_resultChan(strs, &ans, adChan, exitChan)
	}

	for i := 0; i < peoples; i++ {
		<-exitChan
	}
	close(exitChan)
	//返回数据
	for k, v := range ans {
		result = append(result, model.Item{RedisId: this.RedisID, DocId: k, Score: v})
	}

	return
}
