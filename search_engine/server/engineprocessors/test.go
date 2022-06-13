package engineprocessors

import (
	"Search_Engines/search_engines/server/engineserver"
	"Search_Engines/search_engines/server/model"
	"Search_Engines/search_engines/server/storagesystem/process"
	"Search_Engines/search_engines/server/storagesystem/processor"
	"Search_Engines/search_engines/server/untils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	avl "github.com/emirpasic/gods/trees/avltree"
)

var (
	EGS *engineserver.EngineProcessor
)

type Ad_test struct {
}

func (this *Ad_test) docDao_sample_test() {
	//功能完成对UserDao初始化功能
	a := model.Doc{Url: "http://localhost//123.com", Document: "你好呀！！！"}
	b := model.Doc{Url: "http://localhost//1234.com", Document: "哈哈哈"}
	c := model.Doc{Url: "c", Document: "滚吧你"}

	dao.Write_Doc(1, 1, &a)
	dao.Write_Doc(1, 2, &b)
	dao.Write_Doc(2, 1, &c)

	v, _ := dao.Read_ID(1, 1, 2)
	for _, k := range v {
		k.Debug_Show()
	}
	v, _ = dao.Read_ID(2, 1)
	for _, k := range v {
		k.Debug_Show()
	}
	v, err := dao.Read_ID(2, 2, 1)
	if err != nil {
		fmt.Println("不存在")
	}
}

func (this *Ad_test) engineserver_sample_test() {

	egs := &engineserver.EngineProcessor{}
	egs.Start(1, 0)
	t := untils.New(3)
	t.Add(1)
	t.Add(3)
	egs.IndexProcess.Tree.Put("ddd", t)
	egs.IndexProcess.Tree.Put("addd", t)
	//fmt.Println(egs.IndexProcess.Tree)
	err := egs.End()
	if err != nil {
		fmt.Println("ERRR engineserver_sample_test")
		return
	}
	fmt.Println("Accept engineserver_sample_test")
}
func For_Key_Value(tr *avl.Tree) {
	it := tr.Iterator()
	for i := 0; it.Next(); i++ {
		fmt.Println(it.Key(), " --------- ", it.Value())
		if i > 100 {
			break
		}
	}
}

//https://paste.ubuntu.com/p/n8QHqvj9rk/
func (this *Ad_test) engineProcessor_sample_test() {
	// var gou = []model.Doc{
	// 	model.Doc{Url: "http://localhost//123.com", Document: "中国 中国 中国 中国 你好呀！！！"},
	// 	model.Doc{Url: "http://localhost//1234.com", Document: "哈哈哈"},
	// 	model.Doc{Url: "https://gimg1b9ba84f70ddebdda6601a5576d37c50", Document: "美中国沃可视数码裂隙灯,检查眼前节健康状况"},
	// 	model.Doc{Url: "https898f9764847ea7", Document: "欧美夏季ebay连衣裙 气质圆领通勤绑带收腰连衣裙 zc3730"},
	// 	model.Doc{Url: "https://pd99b8e795.png", Document: "曾是名不见经传的王平,为何能够取代魏延,成为蜀汉"},
	// 	model.Doc{Url: "https://gimg2.baiog2.com&app=20f9ecebd71fe6f27643c17486", Document: "女童黄色连衣裙"},
	// 	model.Doc{Url: "https://udailf636b74a804f0768f6944e0", Document: "探访六盘山隧道 犹如海底世界"},
	// 	model.Doc{Url: "https://gi=jpeg?sece37ca6e7", Document: "蚂蚁摄影宝典上册 64个生活场景学参数 入门书籍教程拍摄技巧"},
	// 	model.Doc{Url: "https://pics0.ba055f7c6b&s=8903925D9776C7CE022DBDF903001033", Document: "部编版二年级《沙滩上的童话》老师整理的详细笔记:知识点都在这"},
	// 	model.Doc{Url: "https://gimg2aee21bc4664d57ceff", Document: "中国的大城市中,中国的大城市中中国的大城市中中国的大城市中中国的大城市中还能见到多少河边洗衣服的女人,这样的景致,在如今中国的大城市中,还能见到多少"},
	// }
/*	egs := &engineserver.EngineProcessor{}
	err := egs.Start(1, 0)
	//fmt.Println("key ", egs.IndexProcess.Keys)

	//fmt.Println(egs.IndexProcess.Tree)

	if err != nil {
		fmt.Println("err 125 test", err)
		return
	}
	// for _, v := range gou {
	// 	//fmt.Println(v.Document)

	// 	err := egs.ADD_Doc(&v)
	// 	if err != nil {
	// 		fmt.Println("add 索引faiL")
	// 	}
	// }
	// (egs.IndexProcess.Tree)

	fmt.Println(egs.IndexProcess.Tree)

	items, err := egs.Index_Qry("中国")
	if err != nil {
		fmt.Println("没发现!!! ")
		return
	}
	for _, item := range items {
		docs, _ := dao.Read_File_ID(egs.RedisID, item.DocId)
		for _, doc := range docs {
			fmt.Println(doc)
		}
		fmt.Println(item.Score)
	}

	For_Key_Value(egs.IndexProcess.Tree)
	egs.End()
	fmt.Println("Accept this.engineProcessor_sample_test()	")
	*/
}
func (this *Ad_test) File_System_sample_test() {
	fp := processor.FileOp{}
	str := "切片和数组的类型有什么不一样，我们可以打印一下，就可以知道两者的区别了，开始 数组是容量的"

	args := []string{"1", "1"}
	fp.Write_File(args, []byte(str))

	data, _ := fp.Read_File(args)
	if string(data) == str {
		fmt.Println("Accept File_System_sample_test 001")
	} else {
		fmt.Println("Fail File_System_sample_test  001")
	}
}

func PUTELEM(sum []model.Doc, adChan chan model.Doc) {
	for _, v := range sum {
		adChan <- v
	}
	close(adChan)
}

func OPT(adChan chan model.Doc, exitChan chan bool) {
	//使用for循环
	//var
	for {
		num, ok := <-adChan
		if !ok { //如果取不到值
			break
		}
		EGS.ADD_Doc(&num)
	}
	log.Println("有一个协程取不数据退出")
	//这里不关闭primeChan
	exitChan <- true
}

//用50k的数据集测试
func (this *Ad_test) Load_Test_Wukong50k_release_Gao(TEXT_PATH string) {

	//读取数据集合
	var docs []model.Doc
	jsonFile, err := os.Open(TEXT_PATH)
	if err != nil {
		fmt.Println("error opening json file")
		return
	}
	defer jsonFile.Close()
	jsonData, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		fmt.Println("error reading json file")
		return
	}
	json.Unmarshal(jsonData, &docs)
	EGS = &engineserver.EngineProcessor{}
	err = EGS.Start(1, 0)
	log.Println("开始我的文档数据 ", EGS.IndexProcess.Keys)
	if Egs.IndexProcess.Keys > 1e8 {
		Egs.End()
		return
	}

	if err != nil {
		fmt.Println("err 125 test", err)
		return
	}

	sum := len(docs)
	peoples := 12
	fmt.Println("开始多线程opt ", sum)
	adChan := make(chan model.Doc, 100)
	exitChan := make(chan bool, peoples)
	go PUTELEM(docs, adChan)

	for i := 0; i < peoples; i++ {
		go OPT(adChan, exitChan)
	}

	for i := 0; i < peoples; i++ {
		<-exitChan
	}

	// uids, err := EGS.Index_Qry("中国")
	// if err != nil {
	// 	fmt.Println("173 ")
	// }
	// //For_Key_Value(egs.IndexProcess.Tree)
	// if uids != nil {
	// 	// docs, _ := dao.Read_File_ID(egs.RedisID, uids.List()...)
	// 	// for i := 0; i < len(docs); i++ {
	// 	// 	fmt.Println(docs[i].Document)
	// 	// }
	// } else {
	// 	fmt.Println("Fail 	this.engineProcessor_sample_test()	")
	// }

	//For_Key_Value(egs.IndexProcess.Tree)
	log.Println("操作后我的索引数量 ", EGS.IndexProcess.Keys)

	EGS.End()
	log.Println("func (this *Ad_test) Load_Test_Wukong50k_release_Gao() end	", TEXT_PATH)
}

func (this *Ad_test) Qry_files() {
	fp := process.Fileprocess{}
	PATH := "/home/xxx/src/Search_Engines/FileStorge/wu_kong_shuju/"
	log.Println(PATH)
	list, err := fp.File_List(PATH)
	if err != nil {
		return
	}
	for _, v := range list {
		filepath := PATH + v + ".json"
		log.Println(filepath)
		this.Load_Test_Wukong50k_release_Gao(filepath)
		time.Sleep(time.Second)
		log.Println("完成一个", filepath)
	}
}
func (this *Ad_test) All_test() {
	//this.Load_Test_Wukong50k_release_Gao()

	//fmt.Println("All_test \n")
	// this.docDao_sample_test()
	// this.engineserver_sample_test()
	//	this.engineProcessor_sample_test()
	//	this.File_System_sample_test()
	//this.Load_Test_Wukong50k_release_Gao()
	this.Qry_files()
}
