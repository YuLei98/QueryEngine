package main

import (
	"Search_Engines/search_engines/commom/static"
	"Search_Engines/search_engines/server/engineprocessors"
	controller "Search_Engines/search_engines/wepserver/controllers"
	"net/http"
	"os"
	"runtime"
	"fmt"
	"github.com/gin-gonic/gin"
)

/*
开发步骤
1.注册路由
2.绑定控制函数
3.控制函数绑定commom中request的结构体参数
4.写请求案例。
*/

func main() {

	cpuNum := runtime.NumCPU() //获得当前设备的cpu核心数
	fmt.Println("cpu核心数:", cpuNum)
	runtime.GOMAXPROCS(cpuNum) //设置需要用到的cpu数量
	
	r := gin.Default()
	r.Static("/static", static.WEP_STATIC_FILE_PATH)
	//r.LoadHTMLGlob("view/*")
	r.LoadHTMLFiles("views/login.html", "views/register.html", "views/pageserver.html",
		"views/index.html", "views/Login_Register.html", "views/wu_kong_search.html",
		"views/userfavorites.html") //模板解析

	r.Any("/pageserver", func(c *gin.Context) {
		//http请求状态码
		c.HTML(http.StatusOK, "pageserver.html", gin.H{})
	})

	// r.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "Login_Register.html", gin.H{})
	// })
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", gin.H{})
	})

	r.GET("/wu_kong_search", func(c *gin.Context) {
		c.HTML(http.StatusOK, "wu_kong_search.html", gin.H{})
	})

	r.GET("/userfavorites", func(c *gin.Context) {
		c.HTML(http.StatusOK, "userfavorites.html", gin.H{})
	})

	//api/
	userGroup := r.Group("/api")
	{
		userGroup.PUT("/wu_kong_search/updata", controller.PUT_STR_MES)
		userGroup.PUT("/wu_kong_search/Epigraph_words",controller.PUT_STR_EPIGRAPH_WORD)
		userGroup.GET("/user/search/:username", controller.GET_QRY_STR_MES)
		userGroup.GET("/search/tips/", controller.GET_USER_KEY_WORD_EXPANSION)
	}
	//user
	userGroup2 := r.Group("/user")
	{
		userGroup2.POST("/login", controller.POST_LOGIN_MES)
		userGroup2.POST("/register", controller.POST_RIGISTER_MES)
	}

	//文件夹api
	userGroup3 := r.Group("/dir")
	{
		userGroup3.GET("/get/:username", controller.GET_BOOK_MARKS_MES)
		userGroup3.POST("/add", controller.POST_BOOK_MARKS_MES)
		userGroup3.PUT("/put", controller.PUT_BOOK_MARKS_MES)
		//delete传输bug还未修复这里用post请求
		userGroup3.POST("/del", controller.DELETE_BOOK_MARKS_MES)
	}

	//关机按键
	r.Any("/shoutdown", func(c *gin.Context) {
		engineprocessors.End()
		os.Exit(0)
	})

	r.Run(":8080")
}


