English | [中文](./README-CN.md)
# go-mod-graph-chart
build chart by go mod graph output

## Install

```shell
$ go get -u github.com/PaulXu-cn/go-mod-graph-chart
```

## Usage

```shell
$ cd goProject
$ go mod graph | gmchart
```

The program will start a http server and open the url in default browser.

## Change & Rebuild

If you has changed js code, the front-end project needs to be rebuilt，and then `go install`
```shell
$ npm run build 
$ go install ./gmchart
```

## License

MIT