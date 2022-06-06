package admin

import (
	"github.com/gin-gonic/gin"

	"QueryEngine/middleware"
	api "QueryEngine/router/api"
)

// 注册路由
func Register(router *gin.Engine) {

	// 设置头部解决跨域问题
	router.Use(middleware.Cors())

	router.GET("/api/", api.Welcome)

	router.POST("/api/query", api.Query)

	router.GET("/api/status", api.Status).POST("/api/status", api.Status)

	router.GET("/api/gc", api.Gc).POST("/api/gc", api.Gc)

	router.GET("/api/word/cut", api.WordCut)

	router.GET("/api/dump", api.Dump).POST("/api/dump", api.Dump)

	router.GET("/api/index", api.AddIndex).POST("/api/index", api.AddIndex)

	router.GET("/api/remove", api.RemoveIndex).POST("/api/remove", api.RemoveIndex)

	// 用户注册、登录、注销
	router.POST("/api/register", api.Register)

	router.POST("/api/login", api.Login)

	auth := router.Group("/api/favorite", middleware.AuthCheck())

	auth.GET("username", api.GetUsername)

	// 获取、新增、删除、重命名个人收藏夹
	auth.GET("get_list", api.GetFavoriteList)
	auth.POST("add", api.AddFavoriteList)
	auth.POST("rename", api.RenameFavoriteList)
	auth.POST("delete", api.DeleteFavoriteList)

	// 获取某个收藏夹的所有结果
	auth.POST("get_items", api.GetFavoriteItems)

	// 新增搜索结果到收藏夹
	auth.POST("add_item", api.AddFavoriteItem)

	// 删除收藏夹里的记录
	auth.POST("delete_item", api.DeleteFavoriteItem)
}
