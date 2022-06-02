# QueryEngine

运行后台服务器：

* 将数据集 wukong50k 放到 main.go 同一级目录 dataset/ 下面.

首次运行需加载数据集
```
go run main.go -initdataset true
```

非首次运行
```
go run main.go
```

其他使用方式（查询和添加索引等）见[gofound文档](https://github.com/newpanjing/gofound/blob/main/docs/api.md)

----------------------------------------------

运行前台服务器：

* 进入VUE/文件夹下，运行

```
node app.js
```