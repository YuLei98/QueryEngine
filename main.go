package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"QueryEngine/router/api"
	"QueryEngine/searcher"

	admin "QueryEngine/router/admin"

	"github.com/gin-gonic/gin"
)

func main() {

	var addr string
	flag.StringVar(&addr, "addr", ":5678", "设置监听地址和端口")

	var dataDir string

	//兼容windows
	dir := fmt.Sprintf(".%sdata", string(os.PathSeparator))

	flag.StringVar(&dataDir, "data", dir, "设置数据存储目录")

	var debug bool
	flag.BoolVar(&debug, "debug", false, "设置是否开启调试模式")

	// var isInitDateset bool
	// flag.BoolVar(&isInitDateset, "initdataset", false, "设置是否初始化数据集")

	// 如果设置了该参数，表示加载dataset文件夹下的数据集
	var datasetDir string
	flag.StringVar(&datasetDir, "dataset", "", "设置数据集存储目录")

	flag.Parse()

	if debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	//处理异常
	router.Use(api.Recover)
	err := router.SetTrustedProxies(nil)
	if err != nil {
		return
	}

	//注册api
	admin.Register(router)

	var Engine = &searcher.Engine{
		IndexPath:   dataDir,
		DatasetPath: "./dataset/",
		// InitDateset: isInitDateset,
		DatasetDir: datasetDir,
	}
	option := Engine.GetOptions()

	go Engine.InitOption(option)
	//保存索引到磁盘
	defer Engine.FlushIndex()
	api.SetEngine(Engine)

	log.Println("API url： \t http://" + addr + "/api")

	err = router.Run(addr)
	defer func() {

		if r := recover(); r != nil {

			fmt.Printf("panic: %s\n", r)

		}

		fmt.Println("-- 2 --")

	}()
	fmt.Println("-- 1 --")
	if err != nil {
		fmt.Println("错误", err)
		return
	}
}
