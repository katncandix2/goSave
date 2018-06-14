## goSave
*本项目主要想要使用go 实现对象存储*
```
beta版所实现功能
1.uuid 生成（全局唯一）
2.对外http接口 (暂时可以只监听某一域名固定端口：后期实现路由功能)
3.支持文件预览、上传下载（beta可以仅实现单节点 并对文件大小类型进行限制）
```


```
模块说明
1.BaseTools 共有一些工具类、参数校验http请求等
2.Constants 常量定义 禁止代码中出现 1 ，2 等无意义数字均用常量定义
3.Entity    实体类封装
4.QueueHelper  队列助手、未来添加
5.Storage      文件存储读取
6.TestUnit     测试用例统一放置于此
7.WebServer    路由、http处理接口放置于此 beta 暂时监听特定端口后期手动实现dispathcher
```
```
git 分支命名
[todo]/[name]/[mode]

例如
feature/guruiqin/addHttpServer
fix/guruiqin/fixXXX
```