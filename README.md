<div align=center>
	<img src="web/src/assets/logo-mini.svg" width="150" height="150" />
    <br/>
    <h1>GoFastAdmin</h1>
    <h4>基于Go+Vue开发的权限管理系统</h4>
</div>



##  👻项目简介
- 后端采用Go、Gin、Gorm、Mysql、Redis
- 前端采用Vue3、TypeScript、Element Plus、vite、pinia
- 权限认证使用Jwt，可通过角色控制菜单、按钮权限
- 快速、敏捷、免费、开源，快速搭建属于你的权限管理系统

## ⛱️在线体验
敬请期待

##  🛸项目启动
```bash
# 进入服务端
cd server

# 初始化模块
go mod init

# 下载所有依赖
go mod tidy

# 启动项目
go run main.go
```
```bash
# 进入前端
cd web

# 安装依赖
cnpm install

# 运行项目
cnpm run dev

# 打包发布
cnpm run build
```


##  🎨功能介绍

- 👉  用户管理：系统用户的维护，分配角色、重置密码等。
- 👉  菜单管理：动态维护菜单，包括目录、菜单、按钮。
- 👉  角色管理：对应角色分配菜单、按钮权限。
- 👉  字典管理：系统中产生的状态、类型等进行维护。
- 👉  配置管理：系统参数维护，参数配置项可决定运行机制逻辑。

## 📖帮助文档

* Gin https://gin-gonic.com/zh-cn/docs/
* Gorm https://gorm.io/zh_CN/docs/index.html
* Vue3 https://cn.vuejs.org/
* Element Plus https://element-plus.gitee.io/zh-CN/

通过以上文档，你可以玩转本项目😎，如有问题可从文档中寻找解决办法。

## 💐特别鸣谢

- 👉 vue-next-admin：https://gitee.com/lyt-top/vue-next-admin

##  💌支持一下

如果觉得框架不错，或者已经在使用了，希望你可以去 [Github](https://github.com/xiaofeichuan/go-fast-admin) 或者 [Gitee](https://gitee.com/xiaofeichuan/go-fast-admin) 帮我点个 ⭐ Star，这将是对我极大的鼓励与支持。





