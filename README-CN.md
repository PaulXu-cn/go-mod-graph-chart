[English](./README.md) | 中文

[![go-recipes](https://raw.githubusercontent.com/nikolaydubina/go-recipes/main/badge.svg?raw=true)](https://github.com/nikolaydubina/go-recipes)

# go-mod-graph-chart
一个能将 `go mod graph` 输出内容可视化的无依赖小工具

## 安装

```shell
$ go get -u github.com/PaulXu-cn/go-mod-graph-chart/gmchart
```

Go v1.16 或者更高版本使用如下命令安装

```shell
$ go install github.com/PaulXu-cn/go-mod-graph-chart/gmchart@latest
```

## 使用

```shell
$ cd goProject
$ go mod graph | gmchart
```

执行 `go mod graph` 命令，输出的文本作为该程序的输入，该程序会起一个http服务，并打开 `url` 展示图表

![show](./show.gif)
## 改动重建

如果你改动了 `JS` 代码，记得重新构建前端项目，然后重新构建 `go` 项目
```shell
$ npm run build 
$ go install ./gmchart
```

## 开源协议

MIT
