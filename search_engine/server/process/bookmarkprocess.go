package process

import (
	"Search_Engines/search_engines/commom/message"
	"Search_Engines/search_engines/server/model"
	"Search_Engines/search_engines/server/storagesystem/processor"
	"encoding/json"
	"fmt"
	"math"

	uuid "github.com/satori/go.uuid"
)

/*
管理某一个用户
用户收藏夹
用户赠删改用户收藏夹
*/
type Bookmarkprocess struct {
	Count int      //文档数
	Files []string //文档列表 文件名字
	//uuid 随机生成文件名
}

//第一次初始化从磁盘加载数据该用户收藏夹数据
// 2.开协程优化
func (this *Bookmarkprocess) Init_From_Dir(path string) (err error) {
	fmt.Println("path Init_From_Dir ", path)

	fp := &processor.FileOp{}
	files, err := fp.File_List(path + "/dir")
	this.Count = 0

	if err != nil {
		return
	}
	this.Count = len(files)
	this.Files = files
	return
}

/*
功能：添加一个文件收藏
参数：
实现：
1.uuid获取文件名
2.生成文件路径
3.数据写入文件
4.跟新count 和 files
4.返回状态消息
*/
func (this *Bookmarkprocess) Add_Doc(path string, data []byte) (err error) {
	//1.uuid获取文件名
	u1 := uuid.Must(uuid.NewV4(), nil)
	filename := fmt.Sprintf("%s", u1)

	//2.生成文件路径
	fp := &processor.FileOp{}

	args := make([]string, 2)
	args[0] = path + "/dir"
	args[1] = filename

	//3.数据写入文件

	err = fp.Write_File(args, []byte(data))
	//4.跟新count 和 files 4.返回状态消息
	if err != nil {
		return
	}
	this.Count=this.Count+1
	this.Files = append(this.Files, filename)

	return
}

/*
功能：删除一个文档
参数：文档所在路径 ，文件名
实现：
1.生成文件系统需要的路径
2.调用文件系统
2.移除改文件
3.更新文档数 移除改管理的文档名
返回：状态响应返回状态
*/
func (this *Bookmarkprocess) Del_Doc(path string, filename string) (err error) {
	//1.生成文件系统需要的路径
	args := make([]string, 2)
	args[0] = path + "/dir"
	args[1] = filename
	//2.调用文件系统
	fp := processor.FileOp{}
	err = fp.Del_File(args)
	if err != nil {
		return
	}

	//3.更新文档数 移除改管理的文档名
	this.Count = this.Count - 1
	for i := 0; i < len(this.Files); i++ {
		if this.Files[i] == "local" {
			this.Files = append(this.Files[:i], this.Files[i+1:]...)
		}
	}

	return
}

/*
功能：更新文档
参数：文档路径 文件名 文档数据
实现：
1.拼接文件系统所需的路径
2.调用文件系统写入数据
3.返回消息
*/
func (this *Bookmarkprocess) Update_Doc(path string, filename string, data []byte) (err error) {
	//1.拼接文件系统所需的路径
	args := make([]string, 2)
	args[0] = path + "/dir"
	args[1] = filename

	//2.调用文件系统写入数据
	fp := processor.FileOp{}
	err = fp.Write_File(args, data)
	//3.返回消息
	return
}

/*
功能：响应用户查询请求
参数：页面 和 每页限制的大小
实现：
1. 计算 页面数和每个页面的数据
2. 遍历所需文档
3. 对应每个文档生成对应文件系统的查询路径
4. 从磁盘读取文档格式化成响应消息
5.反序列化成文档 拼接响应消息
返回：响应的数据
*/
func (this *Bookmarkprocess) Response(path string, pageNo int, limit int) *message.ResponseBookMarksDatas {

	fmt.Println(path, " ", pageNo, " ", limit)
	fmt.Println(this.Files)
	fmt.Println(this.Count)

	//计算总页数
	// 1.根据数据集求总页面数
	pagecount := int(math.Ceil(float64(this.Count) / float64(limit)))
	if pagecount > limit {
		return &message.ResponseBookMarksDatas{Count: 0}
	}
	//2. 遍历所需文档
	firstmark := (pageNo - 1) * limit
	cur := firstmark
	response := &message.ResponseBookMarksDatas{Count: this.Count}

	//2. 从磁盘读取文档格式化成响应消息
	fp := &processor.FileOp{}
	args := make([]string, 2)
	args[0] = path + "/dir"
	for i := 0; i < limit; i++ {
		cur = firstmark + i
		if cur >= len(this.Files) {
			break
		}

		args[1] = this.Files[cur]
		data, err := fp.Read_File(args) //数据文url-Caption
		if err != nil {
			continue
		}
		//反序列化成文档 拼接响应消息
		var doc model.Doc
		err = json.Unmarshal(data, &doc)

		if err != nil {
			continue
		}
		reponsedata := message.ResponseBoolMarksData{
			Filename: this.Files[i],
			Url:      doc.Url,
			Caption:  doc.Document,
		}

		response.Datas = append(response.Datas, reponsedata)
	}
	//	fmt.Println("Bookmarkprocess 完成 返回",response)
	//	os.Exit(0)
	return response
}
