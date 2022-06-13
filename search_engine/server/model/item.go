package model

/*

*/
type Item struct {
	RedisId int //数据库ID
	DocId int //文档ID
	Score    float64 //查询的文档得分
}
