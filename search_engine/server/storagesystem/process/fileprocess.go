package process

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Fileprocess struct {
}

//生成路径
func (this *Fileprocess) Make_PATH(path string) {
	if this.Is_Exist(path) == false {
		os.Mkdir(path, os.ModePerm)
	}
	if this.Is_Exist(path) == false {
		os.MkdirAll(path, os.ModePerm)
	}
}

//判断当前路径是否存在
func (this *Fileprocess) Is_Exist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}

// 查询文件基本信息
func (this *Fileprocess) File_Shuju(file string) {
	fi, err := os.Stat(file)
	if err == nil {
		fmt.Println("name:", fi.Name())       //文件名称
		fmt.Println("size:", fi.Size())       //文件大小
		fmt.Println("is dir:", fi.IsDir())    //是否为文件夹
		fmt.Println("mode::", fi.Mode())      //文件权限
		fmt.Println("modTime:", fi.ModTime()) //文件修改时间
	}
}

//查看PATH目录下所有的文件非文件夹
/*
1.如果path不存在则直接返回
*/
func (this *Fileprocess) File_List(path string) (files []string, err error) {
	root := path
	if this.Is_Exist(path) == false {
		return files, errors.New("空文件")
	}
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() == false {
			//fmt.Println(info.Name())
			filenames := strings.Split(info.Name(), ".")
			// fmt.Println(filenames);
			files = append(files, filenames[0])
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return
}

// 删除文件
func (this *Fileprocess) Delete_File(file string) (err error) {
	if this.Is_Exist(file) {
		err = os.Remove(file)
		if err == nil {
			fmt.Println("文件删除成功...")
			return
		} else {
			fmt.Println("文件删除失败")
			return
		}
	}
	fmt.Println("文件未存在....")
	return
}
