English | [中文](./README-CN.md)
# go-mod-graph-chart
A tool build chart by `go mod graph` output with zero dependencies

## Install

```shell
$ go get -u github.com/PaulXu-cn/go-mod-graph-chart/gmchart
```

## Usage

```shell
$ cd goProject
$ go mod graph | gmchart
```

The program will start a http server and open the url in default browser.

![show](./show.gif)

## Change & Rebuild

If you has changed js code, the front-end project needs to be rebuilt，and then `go install`
```shell
$ npm run build 
$ go install ./gmchart
```

## License

MIT