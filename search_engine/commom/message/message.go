package message

/*
定义消息类型和定义与客户端浏览器数据的接口
Datas ====> data
消息设置：
例子：
	GET	/book	查询书籍信息
	POST	/book	创建书籍记录
	PUT	/book	更新书籍信息
	DELETE	/book	删除书籍信息
*/
const (
	DOC_ADD_MES     = "doc_add_mes"     //添加图文对信息
	DOC_ADD_RES_MES = "doc_add_res_mes" //添加图文对返回消息

	DOC_ADD_FAIL = "doc_fail" //图文对添加失败
	DOC_ADD_PASS = "doc_pass" //图文对添加成功

	//登录注册消息
	POST_LOGIN_MES    = "login"        //登录消息
	POST_RIGISTER_MES = "register_mes" //注册消息
)

//对应搜索引擎数据查询
const (
	GET_QRY_STR_MES = "query_string_mes" //对用户进行句子查询
	PUT_STR_MES     = "put_string_mes"   //进行查询句子跟新
)

//定义消息类型
type Message struct {
	Msg   string `json:"type"`  //消息类型
	Datas []byte `json:"datas"` //消息的数据集
}


//收藏夹消息
const(
	GET_BOOK_MARKS_MES="query_book_mes" //查询用户收藏夹
	POST_BOOK_MARKS_MES="post_book_mes" //用户添加收藏
	PUT_BOOK_MARKS_MES="put_book_mes" //编辑用户收藏夹消息
	DELETE_BOOK_MARKS_MES="delete_book_mes" //用户删除收藏夹数据
)

//用户输入关键词提示
const(
	GET_USER_KEY_WORD_EXPANSION="user_key_word_expansion" //用户关键词联想
	PUT_STR_EPIGRAPH_WORD="user_Epigraph_words"//用户敏感词
)
