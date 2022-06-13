package model

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/garyburd/redigo/redis"
)

var (
	indao *IndexEngineDao
)

type IndexEngineDao struct {
	pool *redis.Pool
}

func New_Init_Index_Engine_Dao(pool *redis.Pool) (indexDao *IndexEngineDao) {
	indexDao = &IndexEngineDao{
		pool: pool,
	}
	indao = indexDao
	return
}

func New_IndexEngineDao() (indexDao *IndexEngineDao) {
	return indao
}

/*
功能：根据ID和数据库地址写入文件
参数：redisId:数据库ID docID:文档ID doc:文档内容
实现：
	1.从链接池中取出一根链接
	2.根据数据库ID切换数据库
	3.对result数据压缩// 优化有bug
	4.indexEngine序列化
	5.写入数据库
返回：成功返回nil 失败返回错误
*/
func (this *IndexEngineDao) Write_Index_Engine(redisId int, engineId int, ie *IndexEngine) (err error) {
	//	1.从链接池中取出一根链接
	conn := this.pool.Get()
	defer conn.Close()
	//	2.根据数据库ID切换数据库
	_, err = conn.Do("select", redisId)
	if err != nil {
		log.Println("切换数据库失败!!!", err)
		return
	}

	//4.indexEngine序列化
	data, err := json.Marshal(ie)
	if err != nil {
		log.Println("IndexEngine Fail", err)
		return
	}

	//	5.写入数据库
	_, err = conn.Do("Hset", "engines", engineId, data)
	if err != nil {
		fmt.Println("//	5.写入数据库 fail\n", err)
		return
	}
	return
}

/*
功能：根据引擎key读取对应数据库
参数：redisId:数据库ID 引擎key
实现：
	1.从链接池中取出一根链接
	2.根据数据库ID切换数据库
	3.根据引擎ID查询结果
	4.反序列化得到结果 --》indexEngine
	6.返回结果
返回值：data:字符串 err
*/

func (this *IndexEngineDao) Read_Index_Engine_ID(redisID int, engineId uint64) (ie *IndexEngine, err error) {
	//	1.从链接池中取出一根链接

	conn := this.pool.Get()

	defer conn.Close()
	//	2.根据数据库ID切换数据库
	_, err = conn.Do("select", redisID)
	if err != nil {
		fmt.Println("切换数据库失败!!!", err)
		return
	}
	//	3.根据引擎ID查询结果
	data, err1 := redis.String(conn.Do("Hget", "engines", engineId))
	if err1 != nil {
		fmt.Println("Hget docs Fail \n", err1)
		return nil, err1
	}
	//	4.反序列化得到结果 --》indexEngine
	ie = &IndexEngine{}
	err = json.Unmarshal([]byte(data), ie)
	if err != nil {
		//	fmt.Println("	err=json.Unmarshal([]byte(data),ie)		",err);
		return
	}
	return
}
