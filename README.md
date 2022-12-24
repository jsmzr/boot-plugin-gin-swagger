# boot-plugin-gin-swagger

[![Build Status](https://github.com/jsmzr/boot-plugin-gin-swagger/workflows/Run%20Tests/badge.svg?branch=main)](https://github.com/jsmzr/boot-plugin-gin-swagger/actions?query=branch%3Amain)
[![codecov](https://codecov.io/gh/jsmzr/boot-plugin-gin-swagger/branch/main/graph/badge.svg?token=HNQCAN3UVR)](https://codecov.io/gh/jsmzr/boot-plugin-gin-swagger)

boot 系列 gin swagger 插件

## 使用说明

在源码按照 swag 规范编写好文档后

1. 拉取依赖 `go get -u github.com/jsmzr/boot-plugin-gin-swagger`
2. 在 `main.go` 入口文件中显式声明插件 `import _ "github.com/jsmzr/boot-plugin-gin-swagger"`
3. 本地安装 swag，`go get github.com/swaggo/swag/cmd/swag`
4. 本地使用 swag 生成 swagger 文件 `swag init`
5. 在 `main.go` 入口文件中显式声明文档，如 `import _ "docs"` 

默认使用 `/swagger/index.html` 访问，可通过修改配置来设置路由。

