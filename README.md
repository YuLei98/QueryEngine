# QueryEngine

## 运行方式：
### *运行后台服务器：*

如需加载数据集，可在运行时添加参数dataset
```
go run main.go [-dataset 数据集所在路径]
```

其他使用方式（查询和添加索引等）见[gofound文档](https://github.com/newpanjing/gofound/blob/main/docs/api.md)


### *运行前台服务器：*

* 进入VUE/文件夹下，运行

```
node app.js
```
------------------------------
## 用户管理
### *注册*
| 接口地址 | /api/register |
| ------ | ------ |
| 请求方式 | POST |
| 表单参数| username, password|

### *登录*
| 接口地址 | /api/login |
| ------ | ------ |
| 请求方式 | POST |
| 表单参数| username, password|



------------------------------
## 收藏夹管理

### *获取收藏夹列表*
| 接口地址 | /api/favorite/get_list |
| ------ | ------ |
| 请求方式 | GET |
### *新建收藏夹*
| 接口地址 | /api/favorite/add |
| ------ | ------ |
| 请求方式 | POST |
| 表单参数 | favorite_name|

### *删除收藏夹*
| 接口地址 | /api/favorite/delete |
| ------ | ------ |
| 请求方式 | POST |
| 表单参数 | favorite_name|

### *重命名收藏夹*
| 接口地址 | /api/favorite/rename |
| ------ | ------ |
| 请求方式 | POST |
| 表单参数 | from, to|

### *获取收藏夹里的搜索记录*
| 接口地址 | /api/favorite/get_items |
| ------ | ------ |
| 请求方式 | POST |
| 表单参数 | favorite_name|
### *添加搜索记录*
| 接口地址 | /api/favorite/add_item |
| ------ | ------ |
| 请求方式 | POST |
| 表单参数 | favorite_name, doc_id|

### *删除搜索记录*
| 接口地址 | /api/favorite/delete_item |
| ------ | ------ |
| 请求方式 | POST |
| 表单参数 | favorite_name, doc_id|

