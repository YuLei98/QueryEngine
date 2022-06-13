package static

var (
	LOAD_DICTIONARY      = "/home/xxx/src/Search_Engines/search_engines/commom/static/MyData/dict" //自定义词典路径
	LOAD_MY_FILE_SYSTEM  = "/home/xxx/src/Search_Engines/FileStorge/"                              //存储原文间的地方
	WEP_STATIC_FILE_PATH = "/home/xxx/src/Search_Engines/search_engines/wepserver/static"          //wep服务静态文件路径
	LIMIT_QRY_RESPONSE   = 200000                                                                    //限制返回最多的数据量
)
var (
	MY_REDIS      = []int{0, 1}
	MY_USER_REDIS = 0 //用户所在redis的库
)
