# 项目目录结构

- database:

存放各种数据库连接创建文件

- handler

处理 http 请求和返回，相当于Controller

- logs

项目生成的日志记录保存文件

- middleware

存放自定义中间件

- model

模型定义文件

- repository

存放模型数据操作逻辑的文件，通过 interface 方式，方便注入不同数据库

- service

业务逻辑层代码，其实是 handler 和 repository 的一个桥梁，方便解耦

- router

路由中间件

- utils

自定义工具类

- server.go 

server 的入口文件

- docs 

文档记录文件，记录开发表结构设计或者打包文档

