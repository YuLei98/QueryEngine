package model

import (
	"Search_Engines/search_engines/server/storagesystem/processor"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/garyburd/redigo/redis"
)

//服务启动每个数据库初始化一个docDao实例

type DocDao struct {
	pool *redis.Pool
}

var (
	docdao *DocDao
)

func New_Init_DocDao(pool *redis.Pool) (docDao *DocDao) {
	docDao = &DocDao{
		pool: pool,
	}
	docdao = &DocDao{
		pool: pool,
	}
	return
}
func New_Doc_Dao() (result *DocDao) {

	return docdao
}

/*
功能：根据ID和数据库地址写入文件
参数：redisId:数据库ID docID:文档ID doc:文档内容
实现：
	1.从链接池中取出一根链接
	2.根据数据库ID切换数据库
	3.result=doc序列化
	4.对result数据压缩
	5.写入数据库
返回：成功返回nil 失败返回错误
*/
func (this *DocDao) Write_Doc(redisID int, docID int, doc *Doc) (err error) {
	//fmt.Println("func (this *DocDao) Write_Doc(redisID int, docID int, doc *Doc) (err error) stat")
	//	1.从链接池中取出一根链接
	conn := this.pool.Get()
	defer conn.Close()
	//	2.根据数据库ID切换数据库
	_, err = conn.Do("select", redisID)
	if err != nil {
		fmt.Println("切换数据库失败!!!", err)
		return
	}
	//	2.result=doc序列化
	data, err := json.Marshal(doc)

	if err != nil {
		fmt.Println("data,err:=json.Marshal(doc) fail!!!", err)
		return
	}
	//	4.对result数据压缩 还有bug
	// fp:=process.Compressor{}
	// data=fp.DoZlibCompress(data)

	//	5.写入数据库
	_, err = conn.Do("Hset", "docs", docID, data)
	if err != nil {
		fmt.Println("	//	5.写入数据库 fail\n", err)
		return
	}

	//fmt.Println("func (this *DocDao) Write_Doc(redisID int, docID int, doc *Doc) (err error) end ")

	return
}

/*
功能：根据文档ID读取对应数据库文档
参数：redisId:数据库ID docID文档key
实现：
	1.从链接池中取出一根链接
	2.根据数据库ID切换数据库
	3.根据文档ID查询结果
	4.对result数据解压缩
	5.结果反序列化
	6.返回结果
返回：doc err
*/

func (this *DocDao) Read_ID(redisID int, items ...int) (docs []*Doc, err error) {
	//	1.从链接池中取出一根链接
	conn := this.pool.Get()
	defer conn.Close()
	//	2.根据数据库ID切换数据库
	_, err = conn.Do("select", redisID)
	if err != nil {
		fmt.Println("切换数据库失败!!!", err)
		return
	}
	//fp:=process.Compressor{}

	for _, v := range items {
		fmt.Println(v)
		//Hashes
		result, err1 := redis.String(conn.Do("Hget", "docs", v))
		if err1 != nil {
			fmt.Println("Hget docs Fail \n", result)
			return nil, err1
		}
		data := []byte(result)

		doc := &Doc{}
		err = json.Unmarshal(data, doc)
		docs = append(docs, doc)
		if err != nil {
			fmt.Println("err=err=json.Unmarshal(data,doc", err)
			return
		}
	}
	return
}

/*
功能：根据ID和数据库地址写入磁盘文件
参数：redisId:数据库ID docID:文档ID doc:文档内容
实现：
	1.根据数据库ID切换数据库生成文件名
	2.result=doc序列化
	3.调用磁盘管理器处理
	4.写入文件
返回：成功返回nil 失败返回错误
*/
func (this *DocDao) Write_File_Doc(redisID int, docID int, doc *Doc) (err error) {
	//1.根据数据库ID切换数据库生成文件名
	args := []string{strconv.Itoa(redisID) + "/", strconv.Itoa(docID)}
	// filename := strconv.Itoa(redisID) + "/" + strconv.Itoa(docID)

	//	2.result=doc序列化
	data, err := json.Marshal(doc)
	if err != nil {
		fmt.Println("data,err:=json.Marshal(doc) fail!!!", err)
		return
	}
	//调用磁盘管理西贡处理
	fp := &processor.FileOp{}
	err = fp.Write_File(args, []byte(data))
	return
}

/*
功能：根据数据库ID和文档ID从磁盘读取对应数据库文档
参数：redisId:数据库ID docID文档key
实现：
	1.根据数据库ID切换数据库生成文件名
	2.生成磁盘系统对象
	3.根据生成文件名查询结果
	4.结果反序列化
	.添加返回结果
返回：doc err
*/

func (this *DocDao) Read_File_ID(redisID int, items ...int) (docs []*Doc, err error) {
	//1.根据数据库ID切换数据库生成文件名
	//生成磁盘系统对象
	fp := &processor.FileOp{}
	for _, v := range items {
		//1.根据数据库ID切换数据库生成文件名

		args := []string{strconv.Itoa(redisID) + "/", strconv.Itoa(v)}
		//2.根据生成文件名查询结果
		data, err1 := fp.Read_File(args)
		if err1 != nil {
			fmt.Println("磁盘读取失败 Fail", err1)
			return nil, err1
		}
		doc := &Doc{}
		err = json.Unmarshal(data, doc)
		docs = append(docs, doc)
		if err != nil {
			fmt.Println("err=err=json.Unmarshal(data,doc", err)
			return
		}
	}
	return
}
