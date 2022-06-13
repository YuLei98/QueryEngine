package model

type Node struct{
	Key string `json:"key"`    //Key 
	Value map[int]int `json:"value"` //Value 
}
type IndexEngine struct {
	//确定字段 要保证反序列成功
	Tree    []string `json:"tree"`    //avl 树
	RedisId int    `json:"redisId"` //数据库ID
	Keys    uint64 `json:"keys"`    //该数据库的keys值
}
