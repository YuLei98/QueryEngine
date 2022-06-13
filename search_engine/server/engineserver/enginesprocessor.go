package engineserver

import (
	"Search_Engines/search_engines/commom/message"
	"Search_Engines/search_engines/server/model"
	"Search_Engines/search_engines/server/process"
	"encoding/json"
	"log"
	
	avl "github.com/emirpasic/gods/trees/avltree"
)

/*
 功能 管理每个数据库的索引器：
 1.接受一个消息 进行消息处理
 2.管理索引处理器 加载索引处理器
 3.根据传入数据库ID 分词查询 文档ID
 实现；
 1. 一个数据库都有一个索引docs索引管理器
 2.
*/
type EngineProcessor struct {
	IndexProcess *process.Iddexprocess
	RedisID      int
	EngineID     int

	EpigraphWord map[string]string
}

/*
功能：通过redisId  和 key 生成数据一个索引管理器
参数: redisId key
实现：
1.数据库不存在 则生成一个
2.数据库中存在 返回数据库中的索引管理器
返回：*Iddexprocess 一个索引管理器
*/
func (this *EngineProcessor) Create_Index_Process(redisid int, ipId int) (idp *process.Iddexprocess, err error) { //从文件中加载数据库
	ied := model.New_IndexEngineDao()
	ie, err := ied.Read_Index_Engine_ID(redisid, uint64(ipId))
	if err != nil {
		log.Println("数据库中不存在!!!")
		idp = &process.Iddexprocess{
			Tree:    avl.NewWithStringComparator(),
			RedisID: redisid,
			Keys:    0,
		}
		return idp, nil
	}

	//反序列化生成结果
	idp = &process.Iddexprocess{
		Tree:    avl.NewWithStringComparator(),
		RedisID: ie.RedisId,
		Keys:    ie.Keys,
	}
	// 反向生成key value
	for _, v := range ie.Tree {
		//
		var nodec model.Node
		err = json.Unmarshal([]byte(v), &nodec)
		if err != nil {
			continue
		}
		//	fmt.Println(nodec)
		idp.Tree.Put(nodec.Key, nodec.Value)
	}

	return
}

/*
功能：生成一个引擎管理者
参数：数据库id 引擎ID
实现：
 1.生成或者加载一个索引管理器
 2.生成引擎管理者
返回：
*/
func (this *EngineProcessor) Start(redisid int, ipId int) (err error) {

	this.IndexProcess, err = this.Create_Index_Process(redisid, ipId)
	this.RedisID = redisid
	this.EngineID = ipId
	this.EpigraphWord=make(map[string]string)
	return
}

/*
功能:缓存这个索引器到数据库
参数:
实现:
1.遍历所有key-value对
2.拼model.IndexEngine.Tree
3.拼model.IndexEngine RedisID  EngineID
4.传入 redisId ,EnginId 直接Dao层写入
返回:error (描述信息)
*/
func (this *EngineProcessor) End() (err error) {
	ied := model.New_IndexEngineDao()

	var inde model.IndexEngine
	//1.遍历所有key-value对

	it := this.IndexProcess.Tree.Iterator()
	for i := 0; it.Next(); i++ {
		key, ok := it.Key().(string)
		if !ok {
			continue
		}
		val, ok := it.Value().(map[int]int)
		if !ok {
			continue
		}
		//2.拼model.Node
		node := model.Node{Key: key, Value: val}
		data, err := json.Marshal(node)
		if err != nil {
			continue
		}

		//3.拼model.IndexEngine
		inde.Tree = append(inde.Tree, string(data))
	}
	//4.拼model.IndexEngine RedisID  EngineID
	inde.RedisId = this.IndexProcess.RedisID
	inde.Keys = this.IndexProcess.Keys

	//5.传入 redisId ,EnginId 直接Dao层写入
	ied.Write_Index_Engine(this.RedisID, this.EngineID, &inde)
	log.Println("End 关机成功")
	return
}

/*
 功能：增加图文对
 参数：model.Doc
 实现：直接调用索引处理器处理
*/
func (this *EngineProcessor) ADD_Doc(doc *model.Doc) (err error) {
	err = this.IndexProcess.Add_Doc(doc)
	return
}

/*
功能：查询图文对文档ID
参数：string
实现：直接调用索引处理器处理
返回：直接返回查序到的数据集
*/
func (this *EngineProcessor) Index_Qry(msg *message.RequestUpdataStence) (items []model.Item, err error) {

	items, err = this.IndexProcess.Index_Qry(msg.Sentence,this.EpigraphWord[msg.Username])
	return
}


func (this *EngineProcessor) Updata_epigraph_word(msg *message.RequestEpigraphWords)(err error){
	this.EpigraphWord[msg.Username]=msg.InputWord
	return 
}
/*
功能：用户关键字提示
实现：
*/
func (this *EngineProcessor) Request_Word_KEY_WORD_EXPANSION(msg *message.RequestUserkeywordExpansion) (response *message.ResponseUserkeywordExpansion) {
	items, err := this.Index_Qry(&message.RequestUpdataStence{
		Username:msg.Username,
		Sentence:msg.InputWord,
	})
	if err != nil {
		return
	}
	itp := &process.Itemsprocess{}

	items = itp.Heap_Fast(5, items)

	//获取文件
	dao := &model.DocDao{}
	response = &message.ResponseUserkeywordExpansion{}

	for _, v := range items {
		docs, err := dao.Read_File_ID(v.RedisId, v.DocId)
		if err != nil {
			continue
		}
		response.Datas = append(response.Datas, (*docs[0]).Document)
	}
	return
}
