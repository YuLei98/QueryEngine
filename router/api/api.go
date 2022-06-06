package api

import (
	"net/http"
	"runtime"
	"runtime/debug"
	"strconv"

	"QueryEngine/helper"
	"QueryEngine/models"
	"QueryEngine/router/result"
	"QueryEngine/searcher"
	"QueryEngine/searcher/model"
	"QueryEngine/searcher/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Engine *searcher.Engine

func SetEngine(e *searcher.Engine) {
	Engine = e
}

func Query(c *gin.Context) {

	var request = &model.SearchRequest{}
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(200, result.Error(err.Error()))
		return
	}

	//调用搜索
	//r := Engine.Search(request)
	r := Engine.MultiSearch(request)

	// 跨域
	c.Header("Access-Control-Allow-Origin", "*")

	c.JSON(200, result.Success(r))
}

func Gc(c *gin.Context) {
	runtime.GC()

	c.JSON(200, result.Success(nil))
}

// status 获取服务器状态
func Status(c *gin.Context) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	//索引状态
	index := &map[string]any{
		"size":  Engine.GetIndexSize(),
		"shard": Engine.Shard,
		"queue": len(Engine.AddDocumentWorkerChan),
	}

	memory := map[string]any{
		"alloc":         m.Alloc,
		"total":         m.TotalAlloc,
		"sys":           m.Sys,
		"heap":          m.HeapAlloc,
		"heap_sys":      m.HeapSys,
		"heap_idle":     m.HeapIdle,
		"heap_inuse":    m.HeapInuse,
		"heap_released": m.HeapReleased,
		"heap_objects":  m.HeapObjects,
	}
	system := &map[string]any{
		"os":      runtime.GOOS,
		"arch":    runtime.GOARCH,
		"cores":   runtime.NumCPU(),
		"version": runtime.Version(),
	}

	r := gin.H{
		"memory": memory,
		"system": system,
		"index":  index,
		"status": "ok",
	}
	// 获取服务器状态
	c.JSON(200, result.Success(r))
}

func AddIndex(c *gin.Context) {
	document := &model.IndexDoc{}
	err := c.BindJSON(&document)
	if err != nil {
		c.JSON(200, result.Error(err.Error()))
		return
	}

	go Engine.IndexDocument(*document)

	c.JSON(200, result.Success(nil))
}

// dump 持久化到磁盘
func Dump(c *gin.Context) {
	go Engine.FlushIndex()
	c.JSON(200, result.Success(gin.H{
		"size": Engine.GetIndexSize(),
	}))
}

func WordCut(c *gin.Context) {
	q := c.Query("q")
	r := Engine.WordCut(q)
	c.JSON(200, result.Success(r))

}

func Welcome(c *gin.Context) {
	c.JSON(200, result.Success("welcome"))
}

func RemoveIndex(c *gin.Context) {
	removeIndexModel := &model.RemoveIndexModel{}
	err := c.BindJSON(&removeIndexModel)
	if err != nil {
		c.JSON(200, result.Error(err.Error()))
		return
	}

	err = Engine.RemoveIndex(removeIndexModel.Id)
	if err != nil {
		c.JSON(200, result.Error(err.Error()))
		return
	}
	c.JSON(200, result.Success(nil))
}

//Recover 处理异常
func Recover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			c.JSON(200, result.Error(err.(error).Error()))
		}
		c.Abort()
	}()
	c.Next()
}

func GetUsername(c *gin.Context) {
	u, _ := c.Get("user_claims")
	userClaim := u.(*helper.UserClaims)

	c.JSON(http.StatusOK, result.Success(map[string]string{
		"username": userClaim.Name,
	}))
}

// METHOD: POST
// @username: 用户名字
// @password: 用户密码
func Register(c *gin.Context) {
	username, password := c.PostForm("username"), c.PostForm("password")
	if username == "" || password == "" {
		c.JSON(http.StatusOK, result.Error("用户名或密码为空，请重新输入"))
		return
	}

	var cnt int64
	err := models.DB.Where("name=?", username).Model(new(models.UserInfo)).Count(&cnt).Error

	if err != nil {
		c.JSON(http.StatusOK, result.Error("Get user error: "+err.Error()))
		return
	}

	if cnt > 0 {
		c.JSON(http.StatusOK, result.Error("该用户已被注册"))
		return
	}

	data := &models.UserInfo{
		Name:     username,
		Password: password,
		Identity: helper.GetUUID(),
		// Password: helper.GetMd5(password),
	}
	err = models.DB.Create(data).Error
	if err != nil {
		c.JSON(http.StatusOK, result.Error("Create user error: "+err.Error()))
		return
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, result.Success("用户注册成功"))
}

// METHOD: POST
// @username: 用户名字
// @password: 用户密码
func Login(c *gin.Context) {
	username, password := c.PostForm("username"), c.PostForm("password")
	if username == "" || password == "" {
		c.JSON(http.StatusOK, result.Error("用户名或密码为空，请重新输入"))
		return
	}

	// 密码可改成密文存储
	// password = helper.GetMd5(password)
	data := new(models.UserInfo)
	err := models.DB.Where("name = ? AND password = ? ", username, password).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, result.Error("用户名或密码错误"))
			return
		}
		c.JSON(http.StatusOK, result.Error("Get user error: "+err.Error()))
		return
	}

	token, err := helper.GenerateToken(data.Identity, data.Name)
	if err != nil {
		c.JSON(http.StatusOK, result.Error("Get user error: "+err.Error()))
		return
	}

	// c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, result.Success(
		map[string]string{
			"token": token,
		}),
	)
}

// Method: GET
func GetFavoriteList(c *gin.Context) {
	u, _ := c.Get("user_claims")
	userClaim := u.(*helper.UserClaims)

	var res []models.FavoriteList
	models.DB.Where("identity = ?", userClaim.Identity).Find(&res)

	favoriteList := make([]string, 0)
	for _, fav := range res {
		favoriteList = append(favoriteList, fav.Name)
	}
	c.JSON(http.StatusOK, result.Success(map[string]interface{}{
		"favorite_list": favoriteList,
	}))
}

// METHOD: POST
// @favorite_name: 收藏夹名字
func AddFavoriteList(c *gin.Context) {
	favoriteName := c.PostForm("favorite_name")

	if favoriteName == "" {
		c.JSON(http.StatusOK, result.Error("收藏夹名字不能为空"))
	}

	u, _ := c.Get("user_claims")
	userClaim := u.(*helper.UserClaims)
	identity := userClaim.Identity

	data := models.FavoriteList{
		Identity: identity,
		Name:     favoriteName,
	}
	res := models.DB.Create(&data)
	if res.Error != nil {
		c.JSON(http.StatusOK, result.Error("新建收藏夹错误"))
		return
	}
	c.JSON(http.StatusOK, result.Success("新建收藏夹成功"))
}

// METHOD: POST
// @from: 收藏夹原始名字
// @to:   修改后的收藏夹名字
// 由前端判断请求的合法性
func RenameFavoriteList(c *gin.Context) {
	from, to := c.PostForm("from"), c.PostForm("to")

	u, _ := c.Get("user_claims")
	userClaim := u.(*helper.UserClaims)
	identity := userClaim.Identity

	models.DB.Model(models.FavoriteList{}).Where("identity = ? and name = ?", identity, from).Update("name", to)
	c.JSON(http.StatusOK, result.Success("重命名收藏夹成功"))
}

// METHOD: POST
// @favorite_name: 收藏夹名字
func DeleteFavoriteList(c *gin.Context) {
	favoriteName := c.PostForm("favorite_name")
	if len(favoriteName) == 0 {
		c.JSON(http.StatusForbidden, result.Error("要删除的收藏夹名字为空"))
	}

	u, _ := c.Get("user_claims")
	userClaim := u.(*helper.UserClaims)
	identity := userClaim.Identity

	res := models.DB.Where("identity = ? AND name = ?", identity, favoriteName).Delete(&models.FavoriteList{})
	if res.RowsAffected < 1 {
		c.JSON(http.StatusOK, result.Error("无对应收藏夹"))
		return
	}

	models.DB.Where("identity = ? AND name = ?", identity, favoriteName).Delete(&models.FavoriteInfo{})
	c.JSON(http.StatusOK, result.Success("删除收藏夹成功"))
}

// METHOD: POST
// @favorite_name
func GetFavoriteItems(c *gin.Context) {
	favoriteName := c.PostForm("favorite_name")
	if len(favoriteName) == 0 {
		c.JSON(http.StatusForbidden, result.Error("要删除的收藏夹名字为空"))
	}

	u, _ := c.Get("user_claims")
	userClaim := u.(*helper.UserClaims)

	var res []models.FavoriteInfo
	models.DB.Where("identity = ? and name = ?", userClaim.Identity, favoriteName).Find(&res)

	docs := make([]model.ResponseDoc, 0)
	for _, fav := range res {
		buf := Engine.GetDocById(uint32(fav.DocumentId))
		doc := new(model.ResponseDoc)

		if buf != nil {
			//gob解析
			storageDoc := new(model.StorageIndexDoc)
			utils.Decoder(buf, &storageDoc)
			doc.Document = storageDoc.Document
			text := storageDoc.Text

			doc.Text = text
			doc.Id = uint32(fav.DocumentId)
			docs = append(docs, *doc)
		}
	}
	c.JSON(http.StatusOK, result.Success(docs))
}

// METHOD: POST
// @Favorite_name: 收藏夹名字
// @doc_id: 收藏的文档ID
func AddFavoriteItem(c *gin.Context) {
	favoriteName := c.PostForm("favorite_name")
	docId, _ := strconv.Atoi(c.PostForm("doc_id"))
	if favoriteName == "" || docId < 0 {
		c.JSON(http.StatusOK, result.Success("请输入正确的收藏夹名字和文档ID"))
		return
	}

	u, _ := c.Get("user_claims")
	userClaim := u.(*helper.UserClaims)

	item := models.FavoriteInfo{
		Identity:   userClaim.Identity,
		Name:       favoriteName,
		DocumentId: docId,
	}
	res := models.DB.Create(&item)

	if res.RowsAffected < 1 {
		c.JSON(http.StatusOK, result.Success("该文档已被添加"))
		return
	}
	c.JSON(http.StatusOK, result.Success("添加成功"))
}

// METHOD: POST
// @favorite_name: 收藏夹名字
// @doc_id:	需要删除的文档ID
func DeleteFavoriteItem(c *gin.Context) {
	favoriteName := c.PostForm("favorite_name")
	docId, _ := strconv.Atoi(c.PostForm("doc_id"))
	if favoriteName == "" || docId < 0 {
		c.JSON(http.StatusOK, result.Success("请输入正确的收藏夹名字和文档ID"))
		return
	}

	u, _ := c.Get("user_claims")
	userClaim := u.(*helper.UserClaims)
	identity := userClaim.Identity

	res := models.DB.Where("identity = ? AND name = ? AND doc_id = ?", identity, favoriteName, docId).Delete(&models.FavoriteInfo{})

	if res.Error != nil {
		c.JSON(http.StatusOK, result.Error("删除记录错误"))
		return
	}
	if res.RowsAffected < 1 {
		c.JSON(http.StatusOK, result.Error("无对应记录"))
		return
	}
	c.JSON(http.StatusOK, result.Success("删除成功"))
}
