# ginson 开发脚手架

![issue](https://img.shields.io/github/issues/easonchen147/ginson)
![build](https://img.shields.io/github/workflow/status/easonchen147/ginson/Deploy)
![forks](https://img.shields.io/github/forks/easonchen147/ginson)
![start](https://img.shields.io/github/stars/easonchen147/ginson)
![license](https://img.shields.io/github/license/easonchen147/ginson)

一个go gin api 开发脚手架，定义了不同功能范畴的文件夹路径，方便快速上手使用

## 入门

集成了常用的中间件接入，如Mysql、Redis、Mongo、Kafka等

### 目录说明

目录结构

```
/api        ---api路由
/biz        ---实际业务CRUD
/cfg        ---系统配置
/pkg        ---存放项目业务通用依赖代码
/platform   ---系统级代码，如数据库等
/docs       ---说明文档
```
