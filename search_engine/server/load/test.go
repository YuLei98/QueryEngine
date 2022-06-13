package main

import (
	"Search_Engines/search_engines/server/engineprocessors"
	"fmt"
	"log"
	"os"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	
	logFile, err := os.OpenFile("/home/xxx/src/Search_Engines/search_engines/wepserver/logs/logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetPrefix("[写个项目吧ccsu_f]")
	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ldate)
	
	fp := engineprocessors.Ad_test{}
	log.Println("开机")
	fp.All_test()
}
