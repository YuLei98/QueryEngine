package processor

import (
	"Search_Engines/search_engines/commom/static"
	"Search_Engines/search_engines/server/storagesystem/process"
	"fmt"
)

type FileOp struct {
}

/*
功能：写入数据到文件系统
参数：args []string, args[0] 文件路径 args[1] 文件名 string, data 数据 []byte
实现:
1.根据文件名生成路径
2.调用文件处理器生成当前文件
3.调用文件写入函数处理器
返回
*/
func (this *FileOp) Write_File(args []string, data []byte) (err error) {
	//1.根据文件明生成路径
	//fmt.Println("出发吧狗狗！！")
	filepath := static.LOAD_MY_FILE_SYSTEM + args[0]

	//2.调用文件处理器生成当前文件

	fpc := &process.Fileprocess{}
	fpc.Make_PATH(filepath)

	//fmt.Println("filepath ", filepath)
	//fmt.Println(fpc.Is_Exist(filepath))

	//3.调用文件写入函数处理器

	// cpc := process.Compressor{}
	// err = cpc.WriteFil_eObject(filepath+args[1]+".dp", data)
	filename := filepath + "/" + args[1] + ".dp"

	//fmt.Println("写", args[1])

	fp := process.Gzipcomprocess{}
	fp.Write(data, filename) //注意filePath 与 filename区别 一个文件夹 一个是dp文件

	//fmt.Println("完成写文件 ")

	return
}

/*
功能：从磁盘中读取文件
参数：filename 文件名 string, data 数据 []byte
实现:
1.根据文件名生成路径
2.调用文件读取函数
返回
*/
func (this *FileOp) Read_File(args []string) (data []byte, err error) {
	//1.根据文件名生成路径
	filename := static.LOAD_MY_FILE_SYSTEM + args[0] + "/" + args[1] + ".dp"
	//2.调用文件读取函数
	// cpc := process.Compressor{}
	// data, err = cpc.ReadFile_Object(filepath)
	//fmt.Println("go ", filename)
	fp := process.Gzipcomprocess{}
	data, err = fp.Read(filename)
	if err != nil {
		fmt.Println("读取Fail", filename)
		return
	}
	return
}

/*
功能：删除一个文件
*/
func (this *FileOp) Del_File(args []string) (err error) {
	filename := static.LOAD_MY_FILE_SYSTEM + args[0] + "/" + args[1] + ".dp"
	fp := &process.Fileprocess{}
	fp.Delete_File(filename)
	return
}

/*
功能：查询某个文件夹所有的dp文件
实现:
1.生成路径名
2.调用系统函数获取文件名
*/
func (this *FileOp) File_List(path string) (files []string, err error) {

	filename := static.LOAD_MY_FILE_SYSTEM + path

	fp := &process.Fileprocess{}
	files, err = fp.File_List(filename)
	return
}
