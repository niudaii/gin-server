<h1 align="center">
	gin-server
</h1>

## 基本介绍

基于 gin 简单开发的后端脚手架，前端对应 [antdv-web](https://github.com/niudaii/antdv-web) 和 [vue-element-admin-web](https://github.com/niudaii/vue-element-admin-web)。

## 技术选型

web 框架：gin

orm 框架：gorm

权限控制：casbin

```
// 初始化表数据
entities := []adapter.CasbinRule{
	// admin 权限
	{Ptype: "p", V0: "1", V1: "/api/user/create", V2: "POST"},
	{Ptype: "p", V0: "1", V1: "/api/user/delete", V2: "POST"},
	{Ptype: "p", V0: "1", V1: "/api/user/resetPassword", V2: "POST"},
	{Ptype: "p", V0: "1", V1: "/api/user/changePassword", V2: "POST"},
	{Ptype: "p", V0: "1", V1: "/api/user/info", V2: "GET"},
	{Ptype: "p", V0: "1", V1: "/api/user/menu", V2: "GET"},
	{Ptype: "p", V0: "1", V1: "/api/users", V2: "POST"},

	{Ptype: "p", V0: "1", V1: "/api/operations", V2: "POST"},
	
	// user 权限
	{Ptype: "p", V0: "2", V1: "/api/user/changePassword", V2: "POST"},
	{Ptype: "p", V0: "2", V1: "/api/user/info", V2: "GET"},
	{Ptype: "p", V0: "2", V1: "/api/user/menu", V2: "GET"},
}
```

日志记录：zap

身份验证：jwt

前端动态路由

```
// 初始化表数据
entities := []AuthorityMenu{
	// 管理员权限
	{AuthorityId: "1", MenuId: 1},   //
	{AuthorityId: "1", MenuId: 10},  //
	{AuthorityId: "1", MenuId: 101}, //
	{AuthorityId: "1", MenuId: 102}, //
	{AuthorityId: "1", MenuId: 103}, //

	// 普通用户权限
	{AuthorityId: "2", MenuId: 1}, //
}
```

```
// 初始化表数据
entities := []system.Menu{
	{MenuId: 1, Name: "demo", Path: "/", ParentId: 0, Meta: system.Meta{Title: "栗子", Show: true}, Component: "demo/Demo"},
	{MenuId: 10, Name: "system", ParentId: 0, Meta: system.Meta{Title: "系统设置", Show: true, HideChildren: true}, Component: "system/AdminIndex", Redirect: "/system/user"},
	{MenuId: 101, Name: "user", ParentId: 10, Meta: system.Meta{Title: "用户管理", Show: true}, Component: "system/User"},
	{MenuId: 102, Name: "operation", ParentId: 10, Meta: system.Meta{Title: "日志管理", Show: true}, Component: "system/Operation"},
}
```

## 目录结构

```
├── api						（api层）
├── common.yaml					（common配置文件）
├── config					（配置包）
├── global					（全局对象）
├── initialize					（初始化）
├── middleware					（中间件）
├── model					（模型层）
├── router					（路由层）
├── server.go					（程序入口）
├── server.yaml					（server配置文件）
├── service					（service层）
├── source					（source层）
└── utils					（工具包）
```

## 界面预览

参考前端项目。

## 参考项目

https://github.com/flipped-aurora/gin-vue-admin
